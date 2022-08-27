package main

import (
	"github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver"
)

func main() {
	apiserver.NewApp("apiserver").Run()
}
