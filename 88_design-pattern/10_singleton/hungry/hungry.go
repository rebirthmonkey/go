package hungry

type singleton struct {
	mode string
}

var ins *singleton = &singleton{"on"}

func GetInsOr() *singleton {
	return ins
}