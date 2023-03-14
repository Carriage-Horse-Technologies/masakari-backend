package tasks

// タスク定義

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeDispatchAttack = "dispatch:attack"
)

type AttackServerPayload struct {
	Count int
}

func NewAttackServerTask(count int) (*asynq.Task, error) {
	payload, err := json.Marshal(AttackServerPayload{Count: count})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeDispatchAttack, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func HandleAttackServerTask(ctx context.Context, t *asynq.Task) error {
	var p AttackServerPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Attack the server: count=%d", p.Count)
	Benchmark(p.Count)
	return nil
}
