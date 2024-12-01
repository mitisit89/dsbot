package dsbot

import "strings"

func toMdList(list []string) string {
	var sb strings.Builder
	for _, item := range list {
		sb.WriteString("- ")
		sb.WriteString(item)
		sb.WriteString("\n")
	}
	return sb.String()
}
