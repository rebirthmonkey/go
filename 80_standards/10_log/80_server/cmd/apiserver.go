package main

import (
	"github.com/rebirthmonkey/go/80_standards/10_log/80_server/apiserver"
)

func main() {
	apiserver.NewApp("apiserver").Run()
}
