package monitor

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"local.packages/cmd"
)

var (
	InstanceID = map[string]string{
		"startup":  "7001",
		"shutdown": "7002",
	}
)

func DetectDatetime(instanceID string, targetDate time.Time) []time.Time {
	logName := "System"
	layout := "2006/01/02"
	targetDateStr := targetDate.Format(layout)

	getEventLogCmd := fmt.Sprintf(
		`Get-EventLog -LogName %s -InstanceID %s -After "%s 00:00:00" -Before "%s 23:59:59" | sort -Property TimeGenerated | Format-List`,
		logName, instanceID, targetDateStr, targetDateStr,
	)
	resultByte, err := cmd.InvokeCmd(getEventLogCmd)
	if err != nil {
		log.Fatal(err)
	}
	result := string(resultByte)
	if len(result) == 0 {
		if strings.EqualFold(instanceID, InstanceID["startup"]) {
			fmt.Println("[" + targetDateStr + "]" + " PC was not started.")
		} else if strings.EqualFold(instanceID, InstanceID["shutdown"]) {
			fmt.Println("[" + targetDateStr + "]" + " PC was not shutdowned.")
		}
		return nil
	}

	re := regexp.MustCompile(`TimeGenerated      : \d{4}/\d{2}/\d{2} \d{1,2}:\d{1,2}:\d{1,2}`)
	timeGeneratedList := re.FindAllString(result, -1)
	re = regexp.MustCompile(`\d{4}/\d{2}/\d{2} \d{1,2}:\d{1,2}:\d{1,2}`)
	datetimeStrList := make([]string, len(timeGeneratedList))
	for index, timeGeneratedStr := range timeGeneratedList {
		datetimeStrList[index] = re.FindAllString(timeGeneratedStr, -1)[0]
	}

	layout = "2006/01/02 15:04:05"
	datetime := make([]time.Time, len(datetimeStrList))
	for index, datetimeStr := range datetimeStrList {
		datetime[index], err = time.Parse(layout, datetimeStr)
		if err != nil {
			log.Fatal(err)
		}
	}
	return datetime
}

func DetectStartup(targetDate time.Time) []time.Time {
	return DetectDatetime(InstanceID["startup"], targetDate)
}

func DetectShutdown(targetDate time.Time) []time.Time {
	return DetectDatetime(InstanceID["shutdown"], targetDate)
}
