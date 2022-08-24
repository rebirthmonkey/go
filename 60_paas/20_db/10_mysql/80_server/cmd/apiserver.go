package main

import (
	"github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver"
)

func main() {
	apiserver.NewApp("apiserver").Run()
}
