package inhouse

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/dewadg/go-playground-api/internal/executor"
)

func New(cfgFuncs ...Configurator) executor.Executor {
	cfg := config{}
	for _, f := range cfgFuncs {
		f(&cfg)
	}

	return createExecutor(&cfg)
}

func createExecutor(cfg *config) executor.Executor {
	return func(ctx context.Context, payload executor.ExecutePayload) (executor.ExecuteResult, error) {
		file := []byte(strings.Join(payload.Input, "\n"))

		err := ioutil.WriteFile(cfg.tempDir+"/main.go", file, fs.ModePerm)
		if err != nil {
			return executor.ExecuteResult{}, err
		}

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		cmd.Dir = cfg.tempDir

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return executor.ExecuteResult{}, err
		}
		defer stderr.Close()

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return executor.ExecuteResult{}, err
		}
		defer stdout.Close()

		if err = cmd.Start(); err != nil {
			return executor.ExecuteResult{}, err
		}

		errBytes, err := io.ReadAll(stderr)
		if err != nil {
			return executor.ExecuteResult{}, err
		}

		outputBytes, err := io.ReadAll(stdout)
		if err != nil {
			return executor.ExecuteResult{}, err
		}

		if err = cmd.Wait(); err != nil {
			return executor.ExecuteResult{}, err
		}

		if len(errBytes) > 0 {
			return executor.ExecuteResult{}, errors.New(string(errBytes))
		}

		outputStringLines := strings.Split(string(outputBytes), "\n")
		result := executor.ExecuteResult{
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
