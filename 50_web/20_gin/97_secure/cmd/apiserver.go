package main

import "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver"

func main() {
	apiserver.NewApp("demo-app").Run()
}
