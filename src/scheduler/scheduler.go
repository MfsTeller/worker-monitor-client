package scheduler

import (
	"fmt"
	"log"
	"os"

	"local.packages/cmd"
)

const (
	Run      = "run"
	TaskName = "WorkerMonitor"
	ExecFile = "worker-monitor.exe"
)

func IsTaskScheduled() bool {
	getTaskCmd := fmt.Sprintf(
		`Get-ScheduledTask -TaskName "%s"`,
		TaskName,
	)
	_, err := cmd.InvokeCmd(getTaskCmd)
	if err != nil {
		return false
	}
	return true
}

func RegisterScheduledTask(workDir string) {
	if IsTaskScheduled() {
		fmt.Println("WorkerMonitor task is setupped")
		os.Exit(0)
	}

	// set action
	// $action=New-ScheduledTaskAction -Execute "<executionCommand>" -WorkingDirectory "<workingDirectory>"
	action := fmt.Sprintf(
		`$action=New-ScheduledTaskAction -Execute "%s %s" -WorkingDirectory "%s"; `,
		ExecFile, Run, workDir,
	)

	// set trigger
	// $trigger=New-ScheduledTaskTrigger -Atstartup
	trigger := `$trigger=New-ScheduledTaskTrigger -Atstartup; `

	// set task
	// Register-ScheduledTask -TaskName "WorkerMonitor" -Trigger $trigger -Action $action
	scheduledTask := fmt.Sprintf(
		`Register-ScheduledTask -TaskName "%s" -Trigger $trigger -Action $action`,
		TaskName,
	)
	_, err := cmd.InvokeCmd(action + trigger + scheduledTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Setup succeeded")
}

func UnregisterScheduledTask() {
	if !IsTaskScheduled() {
		fmt.Println("WorkerMonitor task is not setupped")
		os.Exit(0)
	}

	// unsetup task
	// Unregister-ScheduledTask -TaskName "WorkerMonitor -Confirm:$false"
	unscheduledTask := fmt.Sprintf(
		`Unregister-ScheduledTask -TaskName "%s" -Confirm:$false`,
		TaskName,
	)
	_, err := cmd.InvokeCmd(unscheduledTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unsetup succeeded")
}
