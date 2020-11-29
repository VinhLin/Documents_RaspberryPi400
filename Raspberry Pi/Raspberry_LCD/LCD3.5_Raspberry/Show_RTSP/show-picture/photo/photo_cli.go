package photo

import (
	"bytes"
	"context"
	"fmt"
	ptcam "git.dinhviso.com.vn/dvs/dvs-core/proto/camera"
	ptloc "git.dinhviso.com.vn/dvs/dvs-core/proto/location"
	"git.dinhviso.com.vn/dvs/vision-core/consts"
	pttr "git.dinhviso.com.vn/dvs/vision-core/proto/transfer"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/pbnjay/pixfont"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"io"
)

type TransferCli struct {
	pttr.TransferSrvService
}

func NewTransferCli() *TransferCli {
	// create a new service
	service := micro.NewService()

	// DONT'T user Init with FLAG override
	// service.Init()
	cli := service.Client()
	cli.Init(grpc.MaxSendMsgSize(consts.MaxSendMsgSize))

	return &TransferCli{
		pttr.NewTransferSrvService(consts.McTransferService, cli),
	}
}

// SetPhoto using microservice
func (a *TransferCli) SetPhoto(cam *ptcam.CameraLocation) error {
	_, err := a.TransferSrvService.SetPhoto(context.Background(), cam)
	return err
}

// GetLocation using microservice
func (a *TransferCli) GetLocation() (*ptloc.Location, error) {
	return a.TransferSrvService.GetLocation(context.Background(), &pttr.RequestLocation{})
}

func jpegDecode(red io.Reader, loc *ptloc.Location) (*bytes.Buffer, error) {
	img, _, err := image.Decode(red)
	if err != nil {
		return nil, err
	}

	imgRGBA := imageToRGBA(img)
	pixfont.DrawString(imgRGBA, 40, 40, fmt.Sprintf("%.6f-%.6f", loc.Latitude, loc.Longitude), color.White)

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, imgRGBA, nil); err != nil {
		return nil, err
	}

	return buf, nil
}

func imageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}
