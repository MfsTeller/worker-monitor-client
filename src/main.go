package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"local.packages/configloader"
	"local.packages/monitor"
	"local.packages/scheduler"
	"local.packages/timer"
	"local.packages/writer"
)

var (
	NtpServer = "time.windows.com"
	Layout    = map[string]string{
		"datetime": "2006/01/02 15:04:05",
		"date":     "2006/01/02",
	}
	ResultFilePath = "../result"
	TargetDate     time.Time
	IsSetup        bool
	IsUnsetup      bool
	IsGet          bool
	IsPost         bool
)

func parseFlag() {
	TargetDateStr := flag.String("d", "", `Target Date: "YYYY-MM-DD"`)
	isSetup := flag.Bool("setup", false, "Setup Mode Flag: true -> execute setup")
	isUnsetup := flag.Bool("unsetup", false, "Unsetup Mode Flag: true -> execute unsetup")
	isGet := flag.Bool("get", false, "Get Mode Flag: true -> execute GET request")
	isPost := flag.Bool("post", false, "Post Mode Flag: true -> execute POST request")
	flag.Parse()

	// target date setting
	if strings.EqualFold(*TargetDateStr, "") {
		// get current time
		var timerImpl timer.Timer = timer.NewTimer(NtpServer, Layout["datetime"])
		TargetDate = timerImpl.GetCurrentTime()
	} else {
		var err error
		TargetDate, err = time.Parse(Layout["date"], *TargetDateStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// setup mode setting
	IsSetup = *isSetup

	// unsetup mode setting
	IsUnsetup = *isUnsetup

	// get mode setting
	IsGet = *isGet

	// post mode setting
	IsPost = *isPost

}

func main() {
	// flag check
	parseFlag()

	// load config.json
	var configImpl configloader.ConfigLoader = configloader.NewConfigLoader()
	configImpl.Load(`../config/config.json`)

	// setup mode
	if IsSetup {
		scheduler.RegisterScheduledTask(configImpl.GetWorkDir())
		os.Exit(0)
	}

	// unsetup mode
	if IsUnsetup {
		scheduler.UnregisterScheduledTask()
		os.Exit(0)
	}

	// get mode
	if IsGet {
		var clientDataList []writer.ClientData
		var writerImpl writer.Writer = writer.NewWriter(clientDataList)
		respBody := writerImpl.Get(1)
		message := fmt.Sprintf(
			"=== GET client data: client ID = %d",
			configImpl.GetClientID(),
		)
		fmt.Println(message)
		fmt.Println(string(respBody))
		os.Exit(0)
	}

	// get current time
	var timerImpl timer.Timer = timer.NewTimer(NtpServer, Layout["datetime"])
	execDate := timerImpl.GetCurrentTimeFormatted()

	// detect start-up date
	startupDatetimeList := monitor.DetectStartup(TargetDate)
	shutdownDatetimeList := monitor.DetectShutdown(TargetDate)
	fmt.Println("=== Start-up datetime")
	fmt.Println(startupDatetimeList)
	fmt.Println("=== Shutdown datetime")
	fmt.Println(shutdownDatetimeList)

	// generate file content
	clientDataList := make([]writer.ClientData, len(startupDatetimeList))
	for index, s := range startupDatetimeList {
		clientDataList[index].ClientID = configImpl.GetClientID()
		clientDataList[index].Name = configImpl.GetName()
		clientDataList[index].ExecDatetime = execDate

		// convert time.Time list to string list
		clientDataList[index].StartupDatetime = timer.GetTimeFormatted(s, Layout["datetime"])
		if index < len(shutdownDatetimeList) {
			clientDataList[index].ShutdownDatetime = timer.GetTimeFormatted(shutdownDatetimeList[index], Layout["datetime"])
		} else {
			clientDataList[index].ShutdownDatetime = "2000/01/01 00:00:00"
		}
	}

	// write file
	var writerImpl writer.Writer = writer.NewWriter(clientDataList)
	layout := "2006-01-02"
	filename := timer.GetTimeFormatted(TargetDate, layout) + ".json"
	filepath := filepath.Join(ResultFilePath, filename)
	writerImpl.Write(filepath, 0666)

	// post mode
	if IsPost {
		reqBody := writerImpl.GetClientData()
		jsonBytes, err := json.Marshal(reqBody)
		if err != nil {
			log.Fatal(err)
		}
		writerImpl.Post(jsonBytes)
		message := fmt.Sprintf(
			"=== POST client data: client ID = %d",
			configImpl.GetClientID(),
		)
		fmt.Println(message)
		os.Exit(0)
	}

}
