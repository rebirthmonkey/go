package main

import "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver"

func main() {
	apiserver.NewApp("apiserver").Run()
}
