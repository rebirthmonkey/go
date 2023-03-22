package activity

type Activity interface {
	GetName() string
	Execute() error
}
