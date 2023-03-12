package mgr

type BuilderScheme []func(*Manager) error

func (b *BuilderScheme) Register(funcs ...func(*Manager) error) {
	for _, f := range funcs {
		*b = append(*b, f)
	}
}

func (b *BuilderScheme) AddToManager(mgr *Manager) error {
	for _, f := range *b {
		if err := f(mgr); err != nil {
			return err
		}
	}
	return nil
}
