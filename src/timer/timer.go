/*
Copyright 2020 The Worker-Monitor-Client Author.
Licensed under the GNU General Public License v3.0.
    https://github.com/MfsTeller/worker-monitor-client/blob/master/LICENSE
*/
package timer

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

// Timer is a interface for startup/shutdown datetime controll.
type Timer interface {
	GetCurrentTime() time.Time
	GetCurrentTimeFormatted() string
}

// TimerData indicates startup/shutdown datetime on the PC.
type TimerData struct {
	ntpServer   string
	printFormat string
}

// NewTimer creates a timer controller.
func NewTimer(ntpServer string, printFormat string) *TimerData {
	t := new(TimerData)
	t.ntpServer = ntpServer
	t.printFormat = printFormat
	return t
}

// GetCurrentTime obtains current time.
func (u *TimerData) GetCurrentTime() (now time.Time) {
	now, err := ntp.Time(u.ntpServer)
	if err != nil {
		log.Fatal(err)
	}
	// JST time
	location := "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	return now.In(loc)
}

// GetCurrentTimeFormatted obtains current time string with TimerData.
func (u *TimerData) GetCurrentTimeFormatted() string {
	now := u.GetCurrentTime()
	return GetTimeFormatted(now, u.printFormat)
}

// GetTimeFormatted obtains current time string.
func GetTimeFormatted(targetTime time.Time, layout string) string {
	return targetTime.Format(layout)
}

// ParseTime converts string time to time.
func ParseTime(layout, datetimeStr string) time.Time {
	datetime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		log.Fatal(err)
	}
	return datetime
}
