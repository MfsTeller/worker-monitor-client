/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/
package cmd

import (
	"os/exec"
)

// InvokeCmd executes a os command.
func InvokeCmd(cmd string) ([]byte, error) {
	execCmd := exec.Command("powershell", "/c", cmd)
	stdoutStderr, err := execCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return stdoutStderr, err
}
