package resolver

import (
	"context"

	"github.com/dewadg/go-playground-api/internal/executor"
	"github.com/google/uuid"
)

type ExecutionRequest struct {
	Input []string
}

type ExecutionResponse struct {
	Output     []string              `json:"output"`
	ErrorLines *[]ExecutionErrorLine `json:"errorLines"`
}

type ExecutionErrorLine struct {
	Line    int32  `json:"line"`
	Column  int32  `json:"column"`
	Message string `json:"message"`
}

func (r *resolver) Execute(ctx context.Context, args struct{ Payload ExecutionRequest }) (ExecutionResponse, error) {
	payload := executor.ExecutePayload{
		SessionID: uuid.New().String(),
		Input:     args.Payload.Input,
	}

	result, err := executor.Do(ctx, executor.DefaultExecutor, payload)
	if err != nil {
		return ExecutionResponse{}, err
	}

	resp := ExecutionResponse{
		Output: result.Output,
	}
	if len(result.ErrorLines) > 0 {
		errorLines := make([]ExecutionErrorLine, len(result.ErrorLines))
		for i := range result.ErrorLines {
			errorLines[i] = ExecutionErrorLine{
				Line:    int32(result.ErrorLines[i].Line),
				Column:  int32(result.ErrorLines[i].Column),
				Message: result.ErrorLines[i].Message,
			}
		}

		resp.ErrorLines = &errorLines
	}

	return resp, nil
}
