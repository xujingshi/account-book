module account-book

go 1.12

require (
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/google/wire v0.4.0
	go-common v0.3.8
	go.uber.org/automaxprocs v1.2.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	google.golang.org/genproto v0.0.0-20190708153700-3bdd9d9f5532
	google.golang.org/grpc v1.22.0
)

replace go-common => git.bilibili.co/platform/go-common v0.3.75
