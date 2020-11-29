package main

import (
	"flag"
	"os"
	"runtime"

	"git.dinhviso.com.vn/dvs/vision-core/config"
)

var VERSION, GITCOMMIT, DATE string

func main() {
	// setting timezone globally
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Get flags config
	var sconfig string
	flag.StringVar(&sconfig, "config", "/etc/visionclient.toml", "Config file")
	flag.Parse()

	// get variable when build code
	println("Version:", VERSION)
	println("Git:", GITCOMMIT)
	println("Build AT:", DATE)
	println("TimeZone:", os.Getenv("TZ"))
	println("Config:", sconfig)

	// Init config from file toml
	config.Init(sconfig, VERSION, "visionphoto", false)

	runShowPicture()
}
