module github.com/zero-dora/laracom/demo-cli

go 1.14

require (
	github.com/micro/go-micro v1.18.0
	github.com/zero-dora/laracom/demo-service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5
	google.golang.org/grpc v1.25.1
)

replace github.com/zero-dora/laracom/demo-service => F:\project\go\src\laracom\demo-service
