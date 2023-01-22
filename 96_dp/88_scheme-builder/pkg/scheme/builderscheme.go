package scheme

type BuilderScheme []func(*Scheme) error

func (b *BuilderScheme) Register(funcs ...func(*Scheme) error) {
	for _, f := range funcs {
		*b = append(*b, f)
	}
}

func (b *BuilderScheme) AddToManager(mgr *Scheme) error {
	for _, f := range *b {
		if err := f(mgr); err != nil {
			return err
		}
	}
	return nil
}
