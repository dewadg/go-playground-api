package executor

import "context"

type Executor func(ctx context.Context, payload ExecutePayload) (ExecuteResult, error)

type ExecutePayload struct {
	Input []string
}

type ExecuteResult struct {
	Output []string
}

func Do(ctx context.Context, executor Executor, payload ExecutePayload) (ExecuteResult, error) {
	return executor(ctx, payload)
}
