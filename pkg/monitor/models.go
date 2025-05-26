package monitor

type Engine struct {
	state     bool
	startTime string
}

func NewEngine() *Engine {
	return &Engine{
		state:     false,
		startTime: "",
	}
}
