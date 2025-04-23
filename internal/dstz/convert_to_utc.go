package dstz

import (
	"fmt"
	"strings"
	"time"
)

func ToUnixDiscordTimestamp(day string, userTime string, timeZone string) (string, error) {
	layout := "2006-01-02 15:04"
	now := strings.Split(time.Now().Format("2006-01-02"), "-")[:2]
	datetime := fmt.Sprintf("%v-%v-%s %s", now[0], now[1], day, userTime)
	tz, err := findTimeZone(timeZone)
	if err != nil {
		return "", fmt.Errorf("ошибка нахождения временной зоны: %w", err)
	}

	location, err := time.LoadLocation(tz)
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
