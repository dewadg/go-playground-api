package inhouse

import (
	"context"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/dewadg/go-playground-api/internal/executor"
)

func Test_createExecutor(t *testing.T) {
	type args struct {
		cfg     *config
		ctx     context.Context
		payload func() executor.ExecutePayload
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    executor.ExecuteResult
	}{
		{
			name: "success",
			args: args{
				cfg: &config{
					tempDir: "./__tests__",
				},
				ctx: context.Background(),
				payload: func() executor.ExecutePayload {
					file, err := ioutil.ReadFile("./golden/main.example")
					if err != nil {
						t.Fatal(err)
					}

					splitLineFile := strings.Split(string(file), "\n")
					payload := executor.ExecutePayload{
						Input: make([]string, len(splitLineFile)),
					}

					for i, line := range splitLineFile {
						payload.Input[i] = line
					}

					return payload
				},
			},
			wantErr: false,
			want: executor.ExecuteResult{
				Output: []string{
					"Hello, world!",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createExecutor(tt.args.cfg)(tt.args.ctx, tt.args.payload())
			if (err != nil) && !tt.wantErr {
				t.Errorf("createExecutor() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createExecutor() = %v, want %v", got, tt.want)
			}
		})
	}
}
