package git

import "os/exec"

type commandExecutor interface {
	Output() ([]byte, error)
}

// This var gets overwritten in tests
var shellCommandFunc = func(name string, arg ...string) commandExecutor {
	return exec.Command(name, arg...)
}
