package main

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {

	host.Init()
	p := gpioreg.ByName("PWM1_OUT") //GPIO13
	t := time.NewTicker(1 * time.Second)
	for l := gpio.Low; ; l = !l {
		p.Out(l)
		<-t.C
	}
}
