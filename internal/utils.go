package dsbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GenerateEmbed(list []string, title string) *discordgo.MessageEmbed {
	fields := make([]*discordgo.MessageEmbedField, len(list))
	for i, item := range list {
		fields[i] = &discordgo.MessageEmbedField{
			Value:  fmt.Sprintf("* **%s**", item),
			Inline: false,
		}
	}

	return &discordgo.MessageEmbed{
		Title:  title,
		Fields: fields,
	}

}
