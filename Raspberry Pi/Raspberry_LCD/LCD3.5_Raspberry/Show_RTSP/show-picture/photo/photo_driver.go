package photo

import (
	"time"

	ptloc "git.dinhviso.com.vn/dvs/dvs-core/proto/location"
	"git.dinhviso.com.vn/dvs/vision-core/ipcam"
	"gobot.io/x/gobot"
)

// Driver
type Driver struct {
	duration   time.Duration // time for take photo
	name       string
	connection gobot.Connection
}

func NewDriver(a *Adaptor, duration int64) *Driver {
	return &Driver{
		duration:   time.Duration(duration) * time.Second,
		name:       gobot.DefaultName(AdaptorPhoto),
		connection: a,
	}
}

// Name returns the Driver Name
func (d *Driver) Name() string { return d.name }

// SetName sets the Driver Name
func (d *Driver) SetName(n string) { d.name = n }

// Connection returns the Driver Connection
func (d *Driver) Connection() gobot.Connection {
	return d.connection
}

func (d *Driver) adaptor() *Adaptor {
	return d.Connection().(*Adaptor)
}

// Start starts the Driver
func (m *Driver) Start() error {
	return nil
}

// Halt halts the Driver
func (m *Driver) Halt() error {
	return nil
}

// OnTakePhoto
func (d *Driver) TakePhoto() {
	d.adaptor().TakePhoto()
}

func (d *Driver) UpdateLocation() {
	d.adaptor().UpdateLocation()
}

func (d *Driver) GetLocation() *ptloc.Location {
	return d.adaptor().GetLocation()
}

// MountIPCamera
func (d *Driver) MountIPCamera(ipcams map[string]*ipcam.IPCam, isBoot bool) error {
	return d.adaptor().MountIPCamera(ipcams, isBoot)
}

// GetIPCams
func (d *Driver) GetIPCams() map[string]*ipcam.IPCam {
	return d.adaptor().GetIPCams()
}
