package main

import "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver"

func main() {
	apiserver.NewApp("apiserver").Run()
}
