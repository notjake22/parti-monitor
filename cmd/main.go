package main

import "parti-monitor/pkg/monitor"

func main() {
	engine := monitor.NewEngine()
	engine.Run()
}
