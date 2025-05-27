package monitor

import (
	"log"
	"parti-monitor/pkg/discord"
	"parti-monitor/pkg/monitor/parti"
	"time"
)

func (e *Engine) Run() {
	log.Println("Running Parti Monitor Engine...")
	res, err := parti.CheckPublicApi("465731")
	if err != nil {
		log.Fatalf("Error checking public API: %v", err)
		return
	}

	if e.state != res {
		if res {
			e.startTime = time.Now().Format(time.RFC3339)
			err = discord.SendWebhook(e.webhookURL)
			if err != nil {
				log.Printf("Error sending webhook: %v", err)
				return
			}
			log.Println("Start webhook sent successfully!")
		} else {
			err = discord.SendEndWebhook(e.webhookURL, e.startTime, time.Now().Format(time.RFC3339))
			if err != nil {
				log.Printf("Error sending end webhook: %v", err)
				return
			}
			e.startTime = ""
			log.Println("End webhook sent successfully!")
		}

		e.state = res
	}

	time.Sleep(30 * time.Second)
	e.Run()
}
