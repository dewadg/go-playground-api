package executor

import (
	"context"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
)

func Test_createInhouseExecutor(t *testing.T) {
	type args struct {
		cfg     *inhouseConfig
		ctx     context.Context
		payload func() ExecutePayload
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    ExecuteResult
	}{
		{
			name: "success",
			args: args{
				cfg: &inhouseConfig{
					tempDir:      "./__tests__",
					numOfWorkers: 2,
				},
				ctx: context.Background(),
				payload: func() ExecutePayload {
					file, err := ioutil.ReadFile("./golden/main.example")
					if err != nil {
						t.Fatal(err)
					}

					splitLineFile := strings.Split(string(file), "\n")
					payload := ExecutePayload{
						SessionID: "123",
						Input:     make([]string, len(splitLineFile)),
					}

					for i, line := range splitLineFile {
						payload.Input[i] = line
					}

					return payload
				},
			},
			wantErr: false,
			want: ExecuteResult{
				Output: []string{
					"Hello, world!",
				},
				ErrorLines: []ExecuteErrorLine{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createInhouseExecutor(tt.args.cfg)(tt.args.ctx, tt.args.payload())
			if (err != nil) && !tt.wantErr {
				t.Errorf("createExecutor() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createExecutor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_InhouseExecutor(b *testing.B) {
	executor := newInhouse(
		InhouseWithTempDir("./__tests__"),
		InhouseWithNumOfWorkers(4),
	)

	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()

			result, err := executor(context.Background(), ExecutePayload{
				SessionID: uuid.New().String(),
				Input: []string{
					"package main",
					"",
					"import \"fmt\"",
					"",
					"func main() {",
					"fmt.Println(\"Hello, world!\")",
					"}",
					"",
				},
			})
			if err != nil {
				b.Error(err)
				return
			}

			if result.Output[0] != "Hello, world!" {
				b.Errorf("unexpected output: %s", result.Output[0])
			}
		}()
	}

	wg.Wait()
}
