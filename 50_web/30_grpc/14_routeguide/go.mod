module routeguide

go 1.18

replace github.com/rebirthmonkey/pkg/grpc/routeguide => ./

require (
	github.com/golang/protobuf v1.5.3
	github.com/rebirthmonkey/pkg/grpc/routeguide v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.55.0
	google.golang.org/grpc/examples v0.0.0-20230505230727-c44f77e12db9
	google.golang.org/protobuf v1.30.0
)

require (
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)
