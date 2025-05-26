package monitor

import "os"

type Engine struct {
	state      bool
	startTime  string
	webhookURL string
}

func NewEngine() *Engine {
	return &Engine{
		state:      false,
		startTime:  "",
		webhookURL: os.Getenv("WEBHOOK_URL"), // Assuming WEBHOOK_URL is set in the environment
	}
}
