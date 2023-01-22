package main

import (
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/scheme"
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/scheme/registry"

	_ "github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/plugina"
	_ "github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/pluginb"
)

var (
	plugins   = make(map[string]string)
	schemeIns = scheme.Scheme{
		Plugins: plugins,
	}
)

func main() {
	registry.AddToManager(&schemeIns)
	schemeIns.Show()
}
