package resolver

import "context"

type ExecutionResponse struct {
	Output []string `json:"output"`
}

func (r *resolver) Execute(ctx context.Context, args struct{ Content string }) (ExecutionResponse, error) {
	return ExecutionResponse{
		Output: []string{"Hello, world"},
	}, nil
}
