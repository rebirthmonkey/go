package main

import "github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/apiserver"

func main() {
	apiserver.NewApp("demo-app").Run()
}
