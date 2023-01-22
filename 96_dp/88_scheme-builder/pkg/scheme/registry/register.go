package registry

import (
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/scheme"
)

var (
	BuilderScheme scheme.BuilderScheme
)

func Register(funcs ...func(*scheme.Scheme) error) {
	BuilderScheme.Register(funcs...)
}

func AddToManager(mgr *scheme.Scheme) {
	BuilderScheme.AddToManager(mgr)
}
