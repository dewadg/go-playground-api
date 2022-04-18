package executor

import (
	"context"
	"os"
)

type Executor func(ctx context.Context, payload ExecutePayload) (ExecuteResult, error)

type ExecutePayload struct {
	SessionID string
	Input     []string
}

type ExecuteResult struct {
	IsError    bool
	Output     []string
	ErrorLines []ExecuteErrorLine
	Duration   int32
}

type ExecuteErrorLine struct {
	Line    int
	Column  int
	Message string
}

var DefaultExecutor = NewInhouse(
	InhouseWithTempDir(os.Getenv("TEMP_DIR")),
	InhouseWithNumOfWorkers(4),
)

func Do(ctx context.Context, executor Executor, payload ExecutePayload) (ExecuteResult, error) {
	return executor(ctx, payload)
}
