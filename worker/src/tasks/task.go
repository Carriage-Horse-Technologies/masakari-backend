package tasks

// タスク定義

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeDispatchAttack = "dispatch:attack"
)

type AttackServerPayload struct {
	Count int
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
