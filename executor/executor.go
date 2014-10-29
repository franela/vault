package executor

import (
	"log"
	"os/exec"
	"strings"
)

var logger = &logWriter{}

type logWriter struct {
}

func (*logWriter) Write(input []byte) (n int, err error) {
	log.Printf("%s", input)
	return len(input), nil

}

type Executor interface {
	Output(cmd *exec.Cmd) (string, error)
	Run(cmd *exec.Cmd) error
}

type CMDExecutor struct {
}

func (CMDExecutor) Output(cmd *exec.Cmd) (string, error) {

	log.Printf("Running: %s %s\n", cmd.Path, strings.Join(cmd.Args, " "))
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (CMDExecutor) Run(cmd *exec.Cmd) error {

	log.Printf("Running: %s %s\n", cmd.Path, strings.Join(cmd.Args, " "))
	return cmd.Run()

}
