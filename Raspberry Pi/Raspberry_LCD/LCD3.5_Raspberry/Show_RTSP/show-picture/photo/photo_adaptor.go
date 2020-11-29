//TODO when not found ipcamera -> force reboot? using arp-scan , found IP came, but not found when discover -> reboot
//TODO ip route show 2 default route: eth0 (local) wlan0 ( internet)
//TODO auto plugandplay camera

package photo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"strings"
	"sync"
	"time"

	dnumber "git.dinhviso.com.vn/dvs/dvs-core/number"
	ptcam "git.dinhviso.com.vn/dvs/dvs-core/proto/camera"
	ptloc "git.dinhviso.com.vn/dvs/dvs-core/proto/location"
	"git.dinhviso.com.vn/dvs/vision-core/helper"
	"git.dinhviso.com.vn/dvs/vision-core/ipcam"
	"go.uber.org/zap"
	"gobot.io/x/gobot"
)

const (
	AdaptorPhoto = "OnVifIPCAM"
)

var ()

type PhotoPacket struct {
	CameraLocation ptcam.CameraLocation
	DoneChan       chan<- bool
}

type Config struct {
	ID              string // DeviceID
	StoragePath     string
	Ticker          int  // time between 2 take photo action
	IsBachUploading bool // check old task has running
	IsRecord        bool // enable function record between duration
	Cam             ipcam.ConfigurationShort
}

type Adaptor struct {
	name          string // name of adaptor
	config        *Config
	DataChan      chan PhotoPacket
	TranferCli    *TransferCli
	Location      *ptloc.Location
	LocationMutex sync.Mutex
	IPCams        map[string]*ipcam.IPCam // list of IPCam by IP
	IPCamsMutex   sync.Mutex
	isConnect     bool
	Mutex         sync.Mutex
}

// NewAdaptor returns a new Adaptor
func NewAdaptor(cfg *Config) *Adaptor {
	return &Adaptor{
		name:          gobot.DefaultName(AdaptorPhoto),
		config:        cfg,
		IPCams:        make(map[string]*ipcam.IPCam),
		DataChan:      make(chan PhotoPacket),
		TranferCli:    NewTransferCli(),
		LocationMutex: sync.Mutex{},
		IPCamsMutex:   sync.Mutex{},
		isConnect:     false,
		Mutex:         sync.Mutex{},
	}
}

// Name returns the Adaptor Name
func (a *Adaptor) Name() string { return a.name }

// SetName sets the Adaptor Name
func (a *Adaptor) SetName(n string) { a.name = n }

func (a *Adaptor) SetDeviceID(id string) error {
	if dnumber.IsDeviceID(id) == true {
		a.config.ID = id
	} else {
		return fmt.Errorf("Not Valid ID")
	}
	return nil
}

func (a *Adaptor) GetDeviceID() string {
	return a.config.ID
}

func (a *Adaptor) GetPhotoDir(macAddr string) string {
	return path.Join(a.config.StoragePath, strings.Replace(macAddr, ":", "", -1))
}

func (a *Adaptor) GetPhotoDirByDay(macAddr string) string {
	return path.Join(a.GetPhotoDir(macAddr), fmt.Sprintf("%s", time.Now().Format("20060102")))
}

// SendPhoto to NATS server and from it sync dropbox
func (a *Adaptor) SendPhoto() {
	go func() {
		for pkt := range a.DataChan {
			pkt.DoneChan <- true
			close(pkt.DoneChan)

			if err := a.SetPhotoToTransferSrv(&pkt.CameraLocation); err != nil {
				zap.L().Error("SetPhotoToTransferSrv", zap.Error(err))
			}
		}
	}()
}

// Connect establishes a connection to the GPS adaptor
func (a *Adaptor) Connect() (err error) {
	if err := a.MountIPCamera(map[string]*ipcam.IPCam{}, false); err != nil {
		zap.L().Error("MountIPCamera", zap.Error(err))
	}

	return nil
}

// MountIPCamera and UnMount scan and process mount system
func (a *Adaptor) MountIPCamera(ipcams map[string]*ipcam.IPCam, isBoot bool) error {
	fmt.Printf("MountIPCamera IPEXIST:%+v\n", ipcams)
	if len(ipcams) == 0 {
		var err error
		ipcams, err = ipcam.ScanIPCamera(isBoot)
		if err != nil {
			return err
		}
	}

	for _, cam := range ipcams {
		if err := helper.CheckAndMkdir(a.GetPhotoDir(cam.MAC)); err != nil {
			zap.L().Error("CheckAndMkdir", zap.Error(err))
		}

		if err := cam.UpdateStreamUri(1); err != nil {
			zap.L().Error("UpdateSnapshotUri", zap.Error(err))
		}

	}

	a.SetConnected(true)
	a.setIPCams(ipcams)

	return nil
}

// setIPCams newest
func (a *Adaptor) GetIPCams() map[string]*ipcam.IPCam {
	return a.IPCams
}

// setIPCams newest
func (a *Adaptor) setIPCams(ipcams map[string]*ipcam.IPCam) {
	a.IPCamsMutex.Lock()
	defer a.IPCamsMutex.Unlock()
	for _, cam := range ipcams {
		println("SET:", cam.MAC, cam.IP)
	}
	a.IPCams = ipcams
}

// Finalize terminates the connection to the board
func (a *Adaptor) Finalize() error {

	return nil
}

// TakePhoto get photo from ipcamera
// isTransfer send photo to transfer service
func (a *Adaptor) TakePhoto() {
	// client := http.Client{
	// 	Timeout: 2 * time.Second,
	// }

	for _, cam := range a.IPCams {
		//Snapshot
		// resp, err := client.Get(cam.SnapshotURI)
		// if err != nil {
		// 	zap.L().Error("client.Get(cam.SnapshotURI)", zap.String("SnapshotURI", cam.SnapshotURI), zap.Error(err))
		// 	continue
		// }
		// defer resp.Body.Close()

		// if err := a.WriteImage(resp.Body, cam.IP); err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		//Stream
		if a.isConnect {
			if err := a.CommandShowImage(cam.StreamUri); err != nil {
				log.Println(err)
			}
		}

	}
}

// sendToTransfer
// return time at send
func (a *Adaptor) sendToTransfer(mac string, data []byte) uint32 {
	atTime := uint32(time.Now().Unix())
	doneChan := make(chan bool)
	a.DataChan <- PhotoPacket{
		DoneChan: doneChan,
		CameraLocation: ptcam.CameraLocation{
			MAC:        mac,
			Unixtime:   atTime,
			Latitude:   a.Location.Latitude,
			Longitude:  a.Location.Longitude,
			DeviceID:   a.config.ID,
			Data:       data,
			DeliveryAt: atTime,
		},
	}
	isDone := <-doneChan
	zap.L().Info("TakePhoto PhotoChan", zap.Bool("done", isDone))
	return atTime
}

// SetLocationToTransferSrv using microservice
func (a *Adaptor) SetPhotoToTransferSrv(cam *ptcam.CameraLocation) error {
	return a.TranferCli.SetPhoto(cam)
}

func (a *Adaptor) UpdateLocation() {
	loc, err := a.TranferCli.GetLocation()

	if err != nil {
		zap.L().Error("GetLocation", zap.Error(err))
	} else {
		a.LocationMutex.Lock()
		defer a.LocationMutex.Unlock()

		a.Location = loc
	}
}

func (a *Adaptor) GetLocation() *ptloc.Location {
	return a.Location
}

//WriteImage
func (a *Adaptor) WriteImage(resp io.Reader, IP string) error {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(resp); err != nil {
		return fmt.Errorf("ReadFrom: ", err)
	}

	filename := fmt.Sprintf("%s_%v.jpg", IP, time.Now().Unix())

	if err := ioutil.WriteFile(path.Join(a.config.StoragePath, filename), buf.Bytes(), 0755); err != nil {
		return fmt.Errorf("WriteFile: ", err)
	}

	return nil
}

func (a *Adaptor) CommandShowImage(URLstream string) error {
	// omxplayer -o local rtsp://192.168.10.47:554/user=admin_password=tlJwpbo6_channel=1_stream=1.sdp?real_stream
	cmd := exec.Command("omxplayer", "-o", "local", URLstream)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v %v", err, string(output))
	}

	a.SetConnected(false)
	return nil
}

func (a *Adaptor) SetConnected(isConnect bool) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	a.isConnect = isConnect
}
