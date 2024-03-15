package util

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"server/env"

	"github.com/hibiken/asynq"
)

type publisherStruct struct {
	Redis *asynq.Client
}

func NewPublisher() *publisherStruct {
	addr := fmt.Sprintf("%v:%v", env.NewEnvironment().REDIS_HOST, env.NewEnvironment().REDIS_PORT)
	redisClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     addr,
		Username: env.NewEnvironment().REDIS_USERNAME,
		Password: env.NewEnvironment().REDIS_PASSWORD,
		DB:       env.NewEnvironment().REDIS_DATABASE,
	})
	return &publisherStruct{
		Redis: redisClient,
	}
}

func (c *publisherStruct) ResetPassword() {
	defer c.Redis.Close()
	type ResetPasswordStruct struct {
		Email     string `json:"email"`
		Token     string `json:"token"`
		ExpiredAt string `json:"expired_at"`
	}
	payload, _ := json.Marshal(ResetPasswordStruct{Email: "riyaneeng@gmail.com", Token: "token yang sangat rahasia", ExpiredAt: "Jam 5"})
	task := asynq.NewTask("email:reset-password", payload)
	info, err := c.Redis.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(1*time.Minute), asynq.Queue("default"))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func (c *publisherStruct) NotifRegister() {
	defer c.Redis.Close()
	type NotifRegister struct {
		Email string `json:"email"`
	}
	payload, _ := json.Marshal(NotifRegister{Email: "riyaneeng@gmail.com"})
	task := asynq.NewTask("email:notif-register", payload)
	info, err := c.Redis.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(1*time.Minute), asynq.Queue("default"))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func (c *publisherStruct) Register() {
	defer c.Redis.Close()
	type Register struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	payload, _ := json.Marshal(Register{Email: "riyaneeng@gmail.com", Password: "rahasia"})
	task := asynq.NewTask("email:register", payload)
	info, err := c.Redis.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(1*time.Minute), asynq.Queue("default"))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
