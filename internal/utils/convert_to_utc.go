package utils

import (
	"fmt"
	"time"
)

const format = "2006-01-02 15:04"

func ToUnixDiscordTimestamp(timeStr string, month_day string) string {
	datetime := fmt.Sprintf("%d-%s %s", time.Now().Year(), month_day, timeStr)
	parsedTime, err := time.Parse(time.RFC1123, datetime)
	fmt.Println(parsedTime)
	if err != nil {
		fmt.Println("failed to parse datetime: %v", err)
	}
	return fmt.Sprintf("<t:%d:t>", parsedTime.Unix())
}
