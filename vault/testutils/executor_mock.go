package testutils

import (
	"os/exec"
)

type MockCMDExecutor struct {
	Path      string
	Arguments []string
	CMDError  error
	CMDOutput string
}

func (me *MockCMDExecutor) Output(cmd *exec.Cmd) (string, error) {
	me.Path = cmd.Path
	me.Arguments = cmd.Args
	return me.CMDOutput, me.CMDError
}

func (me *MockCMDExecutor) Run(cmd *exec.Cmd) error {
	me.Path = cmd.Path
	me.Arguments = cmd.Args
	return me.CMDError

}
