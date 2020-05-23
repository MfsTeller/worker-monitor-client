package timer

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

type Timer interface {
	GetCurrentTime() time.Time
	GetCurrentTimeFormatted() string
}

type TimerData struct {
	ntpServer   string
	printFormat string
}

func NewTimer(ntpServer string, printFormat string) *TimerData {
	t := new(TimerData)
	t.ntpServer = ntpServer
	t.printFormat = printFormat
	return t
}

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

func (u *TimerData) GetCurrentTimeFormatted() string {
	now := u.GetCurrentTime()
	return GetTimeFormatted(now, u.printFormat)
}

func GetTimeFormatted(targetTime time.Time, layout string) string {
	return targetTime.Format(layout)
}

func ParseTime(layout, datetimeStr string) time.Time {
	datetime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		log.Fatal(err)
	}
	return datetime
}
