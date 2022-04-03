package executor

import (
	"context"
	"io"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"strings"
)

func NewInhouse(cfgFuncs ...InhouseConfigurator) Executor {
	cfg := inhouseConfig{}
	for _, f := range cfgFuncs {
		f(&cfg)
	}

	return createInhouseExecutor(&cfg)
}

func createInhouseExecutor(cfg *inhouseConfig) Executor {
	return func(ctx context.Context, payload ExecutePayload) (ExecuteResult, error) {
		file := []byte(strings.Join(payload.Input, "\n"))

		err := ioutil.WriteFile(cfg.tempDir+"/main.go", file, fs.ModePerm)
		if err != nil {
			return ExecuteResult{}, err
		}

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		cmd.Dir = cfg.tempDir

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return ExecuteResult{}, err
		}
		defer stderr.Close()

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return ExecuteResult{}, err
		}
		defer stdout.Close()

		if err = cmd.Start(); err != nil {
			return ExecuteResult{}, err
		}

		errBytes, err := io.ReadAll(stderr)
		if err != nil {
			return ExecuteResult{}, err
		}

		outputBytes, err := io.ReadAll(stdout)
		if err != nil {
			return ExecuteResult{}, err
		}

		_ = cmd.Wait()

		if len(errBytes) > 0 {
			errLines := strings.Split(string(errBytes), "\n")

			return ExecuteResult{
				IsError: true,
				Output:  errLines,
			}, nil
		}

		outputStringLines := strings.Split(string(outputBytes), "\n")
		result := ExecuteResult{
			Output: make([]string, 0),
		}

		for _, line := range outputStringLines {
			if len(line) == 0 {
				continue
			}

			result.Output = append(result.Output, line)
		}

		return result, nil
	}
}
