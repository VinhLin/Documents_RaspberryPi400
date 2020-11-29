package main

import (
	"path"
	"time"

	"git.dinhviso.com.vn/dvs/vision-core/consts"
	"git.dinhviso.com.vn/dvs/vision-core/ipcam"
	"git.dinhviso.com.vn/show-picture/photo"
	"github.com/jpillora/backoff"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gobot.io/x/gobot"
)

func runShowPicture() {
	var (
		storage  string
		duration int64
		deviceID = viper.GetString("deviceID")
		camCfg   = ipcam.ConfigurationShort{}
	)

	// STORAGE Video
	if viper.IsSet("storages.media") == false {
		storage = path.Join(consts.DefaultStoreDir, deviceID, consts.DefaultPhotoDir)
	} else {
		storage = path.Join(viper.GetString("storages.media"), deviceID, consts.DefaultPhotoDir)
	}

	// Get and set duration takephoto
	if viper.IsSet("photo.duration") {
		duration = viper.GetInt64("photo.duration")
	} else {
		duration = 5
	}

	photoAdaptor := photo.NewAdaptor(&photo.Config{
		ID:          deviceID,
		StoragePath: storage,
		Cam:         camCfg,
	})

	photoDriver := photo.NewDriver(photoAdaptor, duration)

	workPhoto := func() {
		photoDriver.TakePhoto()

		gobot.Every(time.Duration(duration)*time.Second, func() {
			photoDriver.TakePhoto()
		})

		var isBOTickerRun bool

		backoff := &backoff.Backoff{
			Min:    5 * time.Second,
			Max:    60 * time.Second,
			Factor: 2,
			Jitter: true,
		}

		boTick := time.Tick(backoff.Duration())
		go func() {
			for {
				select {
				case <-boTick:
					boTick = time.Tick(backoff.Duration())

					if isBOTickerRun {
						return
					}
					isBOTickerRun = true

					// Scanning
					if ipcams, err := ipcam.ScanIPCamera(false); err != nil {
						zap.L().Error("ScanIPCamera", zap.Error(err))
					} else {
						if len(ipcams) != len(photoDriver.GetIPCams()) {
							if err := photoDriver.MountIPCamera(ipcams, false); err != nil {
								zap.L().Error("MountIPCamera", zap.Error(err))
							}
						}
					}

					isBOTickerRun = false
				}
			}
		}()

	}

	robot := gobot.NewRobot("VisionPictureBot",
		[]gobot.Connection{photoAdaptor},
		[]gobot.Device{},
		workPhoto,
	)

	if err := robot.Start(); err != nil {
		zap.L().Fatal("robot.Start()", zap.Error(err))
	}
}
