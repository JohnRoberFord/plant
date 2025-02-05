package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/JohnRobertFord/go-plant/internal/agent"
)

func main() {

	cfg, err := agent.InitConfig()
	if err != nil {
		log.Fatalf("can't init config: %e", err)
	}
	log.Println(cfg)

	myM := agent.NewStorage()

	ctx := context.Background()
	var wg sync.WaitGroup

	agent.Worker(ctx, func() error {
		myM.CollectMetrics()
		return nil
	}, time.Duration(cfg.PollInterval)*time.Second, 1, &wg)

	agent.Worker(ctx, func() error {
		myM.CollectGopsutilMetrics()
		return nil
	}, time.Duration(cfg.PollInterval)*time.Second, 1, &wg)

	agent.Worker(ctx, func() error {
		agent.SendJSONData(cfg, myM.GetElements())
		myM.ResetCounter("PollCount")
		return nil
	}, time.Duration(cfg.ReportInterval)*time.Second, cfg.RateLimit, &wg)

	wg.Wait()
}
