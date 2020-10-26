package utils

import (
	"github.com/xeonx/timeago"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct{}

type TimestampInterface interface {
	ParseStringToTime(timeString string) (time.Time, error)
	ParseTimeToString(detailTime time.Time) string
	PrettyTime(detailTime time.Time) string
}

func (t *Timestamp) ParseStringToTime(timeString string) (time.Time, error) {
	splitAll := strings.Split(timeString, " ")
	splitDate := strings.Split(splitAll[0], "-")
	splitClock := strings.Split(splitAll[1], ":")
	resYear, err := strconv.Atoi(splitDate[0])
	resMon, err := strconv.Atoi(splitDate[1])
	resDay, err := strconv.Atoi(splitDate[2])
	resHour, err := strconv.Atoi(splitClock[0])
	resMin, err := strconv.Atoi(splitClock[1])
	resSec, err := strconv.ParseFloat(splitClock[2], 64)
	return time.Date(
		resYear,
		time.Month(resMon),
		resDay,
		resHour,
		resMin,
		int(resSec),
		0,
		time.Local), err
}

func (t *Timestamp) ParseTimeToString(detailTime time.Time) string {
	return detailTime.String()
}

func (t *Timestamp) PrettyTime(detailTime time.Time) string {
	return timeago.English.Format(time.Date(
		detailTime.Year(),
		detailTime.Month(),
		detailTime.Day(),
		detailTime.Hour(),
		detailTime.Minute(),
		detailTime.Second(),
		detailTime.Nanosecond(),
		time.Local))
}
