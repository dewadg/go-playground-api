package executor

import (
	"context"
	"io"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
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

		outputChan := fanInStringChan(
			toStringChan(stderr),
			toStringChan(stdout),
		)

		_ = cmd.Wait()

		result := ExecuteResult{
			Output: make([]string, 0),
		}
		for line := range outputChan {
			if len(line) == 0 {
				continue
			}

			result.Output = append(result.Output, line)
		}

		return result, nil
	}
}

func toStringChan(reader io.Reader) chan string {
	outputChan := make(chan string)

	go func() {
		for {
			p := make([]byte, 0, 512)
			n, err := reader.Read(p[:cap(p)])
			if err != nil {
				close(outputChan)
				return
			}

			outputChan <- string(p[:n-1])
		}
	}()

	return outputChan
}

func fanInStringChan(stringChans ...chan string) chan string {
	outputChan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(len(stringChans))

	for _, stringChan := range stringChans {
		go func(input chan string) {
			for line := range input {
				outputChan <- line
			}

			wg.Done()
		}(stringChan)
	}

	go func() {
		wg.Wait()

		close(outputChan)
	}()

	return outputChan
}
