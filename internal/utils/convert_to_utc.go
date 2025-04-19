package utils

import (
	"fmt"
	"time"
)

func ToUnixDiscordTimestamp(day string, userTime string, timeZone string) (string, error) {
	layout := "2006-01-02 15:04"
	now := time.Now()
	datetime := fmt.Sprintf("%v-%v-%s %s", now.Year(), int(now.Month()), day, userTime)
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", fmt.Errorf("error loading location %w", err)
	}

	parsedTime, err := time.ParseInLocation(layout, datetime, location)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга времени: %w", err)
	}
	timestamp := parsedTime.Unix()
	return fmt.Sprintf("<t:%d:f>", timestamp), nil

}
