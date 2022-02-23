package main

import (
	"os"

	"github.com/hibiken/asynq"
	"github.com/nemonicgod/terraforms-api/backend"
	"github.com/nemonicgod/terraforms-api/config"
	tasks "github.com/nemonicgod/terraforms-api/interfaces/scrapers"
)

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

	que := make(map[string]int)
	mux := asynq.NewServeMux()

	wrkr := os.Getenv("WORKER")

	// Setting up which job is put into the queue for this worker
	if wrkr == config.Official {
		que["official"] = 1
		mux.HandleFunc(tasks.TypeOfficial, tasks.HandleOfficialTask)
	}

	b.L.Printf("WRK[%s]: main() starting...", wrkr)

	wrk := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")},
		asynq.Config{
			// If set to a zero or negative value, NewServer will overwrite the value
			// to the number of CPUs usable by the currennt process.
			Concurrency: 10,
			Queues:      que,
			Logger:      b.L,
		},
	)

	if err := wrk.Run(mux); err != nil {
		panic(err)
	}
}
