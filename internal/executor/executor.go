package executor

import (
	"context"
	"os"
)

type executor func(ctx context.Context, payload ExecutePayload) (ExecuteResult, error)

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

var defaultExecutor = newInhouse(
	InhouseWithTempDir(os.Getenv("TEMP_DIR")),
	InhouseWithNumOfWorkers(2),
)

func Do(ctx context.Context, payload ExecutePayload) (ExecuteResult, error) {
	return defaultExecutor(ctx, payload)
}
