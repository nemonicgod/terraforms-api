package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/hibiken/asynq"
	"github.com/nemonicgod/terraforms-api/backend"
	tasks "github.com/nemonicgod/terraforms-api/interfaces/scrapers"
	"github.com/sirupsen/logrus"
)

const (
	QueueOfficial = "official"
)

func logEnqueue(l *logrus.Logger, id, queue string) {
	l.Printf("Client enqueued task: id: %s queue: %s", id, queue)
}

func main() {
	// Initalize a new client, the base entrpy point to the application code
	// the true, true is a poor way of turning DB/Redis on/off, cmon bro
	b, e := backend.NewBackend(true, false)
	if e != nil {
		panic(e)
	}

	// Database connect, defer close
	pqDB, err := b.R.D.DB()
	if err != nil {
		panic(err)
	}
	defer pqDB.Close()

	// Initialize the client for adding tasks to the queue
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
	defer client.Close()

	b.L.Println("SCH: main() starting...")

	// The main scheduler
	sch := gocron.NewScheduler(time.UTC)

	kdx := 0
	sch.Every(1).Day().Do(
		func() {
			task, err := tasks.NewOfficialTask(fmt.Sprintf("job:%d:official", kdx))
			if err != nil {
				panic(err)
			}

			info, err := client.Enqueue(task, asynq.Queue(QueueOfficial), asynq.MaxRetry(0), asynq.Timeout(5*time.Minute))
			if err != nil {
				panic(err)
			}

			logEnqueue(b.L, info.ID, info.Queue)
			kdx++
		},
	)

	sch.StartBlocking()
}
