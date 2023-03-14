package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	"notchman.tech/masakari-worker/tasks"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var redis = asynq.RedisClientOpt{
	Addr: getEnv("REDIS_HOST", "redis:6379"),

	// Use a dedicated db number for asynq.
	// By default, Redis offers 16 databases (0..15)
	DB: 0,
}

func main() {
	// Create and configuring Redis connection.
	redisConnection := asynq.RedisClientOpt{
		Addr: getEnv("REDIS_HOST", "redis:6379"), // Redis server address
	}

	// Create and configuring Asynq worker server.
	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()

	// Define a task handler for the welcome email task.
	mux.HandleFunc(
		tasks.TypeDispatchAttack,     // task type
		tasks.HandleAttackServerTask, // handler function
	)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
