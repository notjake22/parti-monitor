package monitor

import (
	"log"
	"parti-monitor/pkg/discord"
	"parti-monitor/pkg/monitor/parti"
	"time"
)

func (e *Engine) Run() {
	log.Println("Starting Parti Monitor Engine...")
	res, err := parti.CheckPublicApi("465731")
	if err != nil {
		log.Fatalf("Error checking public API: %v", err)
		return
	}

	if e.state != res {
		if res {
			err = discord.SendWebhook("https://discord.com/api/webhooks/1376686371904946207/nFynvA2Mkc6RfXm54IntGNTvnIB_rxJViOXFIyOFO52qG9IqTE65ZOxFc6wa4ORefdV4")
			if err != nil {
				log.Printf("Error sending webhook: %v", err)
				return
			}
		}

		e.startTime = time.Now().Format(time.RFC3339)
		e.state = res
	}

	time.Sleep(60 * time.Second)
	e.Run()
}
