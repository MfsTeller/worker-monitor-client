package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CreateDirectory(dirPath string) {
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatal(err)
	}
}

func DeleteDirectory(dirPath string) {
	if err := os.RemoveAll(dirPath); err != nil {
		log.Fatal(err)
	}
}

func InvokeCmd(cmd string) {
	execCmd := exec.Command("sh", "-c", cmd)
	stdoutStderr, err := execCmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdoutStderr)
}
