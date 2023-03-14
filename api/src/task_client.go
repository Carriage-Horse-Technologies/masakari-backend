package main

import (
	"log"

	"github.com/hibiken/asynq"
	"notchman.tech/masakari-backend/tasks"
)

func createAttackServerTask(count int) {
	redisAddr := getEnv("REDIS_HOST", "localhost:6379")

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately.
	//            Use (*Client).Enqueue method.
	// ------------------------------------------------------

	task, err := tasks.NewAttackServerTask(count)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ------------------------------------------------------------
	// Example 2: Schedule task to be processed in the future.
	//            Use ProcessIn or ProcessAt option.
	// ------------------------------------------------------------

	// info, err = client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	// if err != nil {
	// 	log.Fatalf("could not schedule task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

}
