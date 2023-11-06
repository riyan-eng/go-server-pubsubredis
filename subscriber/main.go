package main

import (
	"fmt"
	"log"
	"subscriber/env"
	"subscriber/util"

	"github.com/hibiken/asynq"
)

func init() {
	env.NewEnvironment()
}

func main() {
	addr := fmt.Sprintf("%v:%v", env.ENV.Redis.Host, env.ENV.Redis.Port)
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     addr,
			Username: env.ENV.Redis.Username,
			Password: env.ENV.Redis.Password,
			DB:       env.ENV.Redis.DB,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:reset", util.NewEmail().ResetPassword)
	mux.HandleFunc("email:register", util.NewEmail().Register)
	mux.HandleFunc("email:inforegister", util.NewEmail().Register)
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
