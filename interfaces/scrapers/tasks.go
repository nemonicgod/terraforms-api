package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/nemonicgod/terraforms-api/config"
	official "github.com/nemonicgod/terraforms-api/interfaces/scrapers/official"
)

var log = config.LoadLoggerGeneric()

const (
	RedisAddr    = "redis:6379"
	TypeOfficial = "scraper:official"
)

type OfficialPayload struct {
	Source string
}

func NewOfficialTask(source string) (*asynq.Task, error) {
	payload, err := json.Marshal(OfficialPayload{Source: source})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeOfficial, payload), nil
}

func HandleOfficialTask(ctx context.Context, t *asynq.Task) error {
	var p OfficialPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("[official] scraper starting... [%s]", p.Source)

	c := make(chan error, 1)
	go func() {
		c <- official.Run()
	}()

	select {
	case <-ctx.Done():
		log.Printf("[official] scraper errored [%s]", p.Source)
		return ctx.Err()
	case res := <-c:
		log.Printf("[official] scraper completed [%s]", p.Source)
		return res
	}
}
