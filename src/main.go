/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/
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

	"local.packages/clientdata"
	"local.packages/configloader"
	"local.packages/monitor"
	"local.packages/scheduler"
	"local.packages/timer"
)

var (
	// TargetDate command line argument parameter
	TargetDate time.Time
	// Layout constant map for print format
	Layout = map[string]string{
		"datetime": "2006/01/02 15:04:05",
		"date":     "2006/01/02",
	}
	// configImpl config file data
	configImpl configloader.ConfigLoader
)

const (
	Setup          = "setup"
	Unsetup        = "unsetup"
	Get            = "get"
	Post           = "post"
	Run            = "run"
	NtpServer      = "time.windows.com"
	ResultFilePath = "../result"
)

// parseFlag configures parameter indicated in command line argument.
func parseFlag() {
	TargetDateStr := flag.String("d", "", `Target Date: "YYYY-MM-DD"`)
	flag.Parse()

	// target date setting
	if strings.EqualFold(*TargetDateStr, "") {
		var timerImpl timer.Timer = timer.NewTimer(NtpServer, Layout["datetime"])
		TargetDate = timerImpl.GetCurrentTime()
	} else {
		TargetDate = timer.ParseTime(Layout["date"], *TargetDateStr)
	}
}

// osArgsValidation validates command line arguments
func osArgsValidation() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments")
		os.Exit(1)
	}
}

func setup() {
	scheduler.RegisterScheduledTask(configImpl.GetWorkDir())
}

func unsetup() {
	scheduler.UnregisterScheduledTask()
}

func get() {
	var clientDataList []clientdata.ClientData
	var clientdataImpl clientdata.ClientDataInterface = clientdata.NewClientData(clientDataList)
	respBody := clientdataImpl.Get(1)
	message := fmt.Sprintf(
		"=== GET client data: client ID = %d",
		configImpl.GetClientID(),
	)
	fmt.Println(message)
	fmt.Println(string(respBody))
}

func run() {
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
	clientDataList := make([]clientdata.ClientData, len(startupDatetimeList))
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
	var clientdataImpl clientdata.ClientDataInterface = clientdata.NewClientData(clientDataList)
	layout := "2006-01-02"
	filename := timer.GetTimeFormatted(TargetDate, layout) + ".json"
	filepath := filepath.Join(ResultFilePath, filename)
	clientdataImpl.Write(filepath, 0666)
}

func post() {
	var clientdataImpl clientdata.ClientDataInterface = clientdata.NewClientData(nil)
	layout := "2006-01-02"
	filename := timer.GetTimeFormatted(TargetDate, layout) + ".json"
	filepath := filepath.Join(ResultFilePath, filename)
	clientdataImpl.Read(filepath)
	reqBody := clientdataImpl.GetClientData()

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	clientdataImpl.Post(jsonBytes)
	message := fmt.Sprintf(
		"=== POST client data: client ID = %d",
		configImpl.GetClientID(),
	)
	fmt.Println(message)
}

func main() {
	parseFlag()

	configImpl = configloader.NewConfigLoader()
	configImpl.Load(`../config/config.json`)

	osArgsValidation()

	targetCmd := os.Args[1]
	switch targetCmd {
	case Setup:
		setup()
	case Unsetup:
		unsetup()
	case Get:
		get()
	case Post:
		post()
	case Run:
		run()
	default:
		fmt.Println("Invalid arguments")
	}
}
