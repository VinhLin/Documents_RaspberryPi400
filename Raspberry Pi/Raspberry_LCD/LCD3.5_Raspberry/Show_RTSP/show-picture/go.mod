module git.dinhviso.com.vn/show-picture

go 1.15

replace git.dinhviso.com.vn/dvs/dvs-core => /home/ducvinh/go/src/git.dinhviso.com.vn/dvs/dvs-core

replace git.dinhviso.com.vn/dvs/vision-core => /home/ducvinh/go/src/git.dinhviso.com.vn/dvs/vision-core

require (
	git.dinhviso.com.vn/dvs/dvs-core v1.0.5
	git.dinhviso.com.vn/dvs/vision-core v0.0.0-00010101000000-000000000000
	github.com/jpillora/backoff v1.0.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pbnjay/pixfont v0.0.0-20200714042608-33b744692567
	github.com/spf13/viper v1.7.1
	github.com/use-go/onvif v0.0.0-20200817103923-4e696ec65aa9
	go.uber.org/zap v1.16.0
	gobot.io/x/gobot v1.14.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible