package cmd

import (
	"os/exec"
)

func InvokeCmd(cmd string) ([]byte, error) {
	execCmd := exec.Command("powershell", "/c", cmd)
	stdoutStderr, err := execCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return stdoutStderr, err
}
