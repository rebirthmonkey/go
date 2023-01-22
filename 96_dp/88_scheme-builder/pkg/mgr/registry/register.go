package registry

import (
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/mgr"
)

var (
	BuilderScheme mgr.BuilderScheme
)

func Register(funcs ...func(*mgr.Manager) error) {
	BuilderScheme.Register(funcs...)
}

func AddToManager(mgr *mgr.Manager) {
	BuilderScheme.AddToManager(mgr)
}
