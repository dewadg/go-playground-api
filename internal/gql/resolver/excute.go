package resolver

import (
	"context"

	"github.com/dewadg/go-playground-api/internal/executor"
)

type ExecutionRequest struct {
	Input []string
}

type ExecutionResponse struct {
	Output []string `json:"output"`
}

func (r *resolver) Execute(ctx context.Context, args struct{ Payload ExecutionRequest }) (ExecutionResponse, error) {
	payload := executor.ExecutePayload{
		Input: args.Payload.Input,
	}

	result, err := executor.Do(ctx, executor.DefaultExecutor, payload)
	if err != nil {
		return ExecutionResponse{}, err
	}

	return ExecutionResponse{
		Output: result.Output,
	}, nil
}
