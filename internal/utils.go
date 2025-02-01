package internal

import (
	"dsbot/internal/storage"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GenerateEmbed(movies []storage.Movie, title string) *discordgo.MessageEmbed {
	fields := make([]*discordgo.MessageEmbedField, len(movies))
	for i, item := range movies {
		fields[i] = &discordgo.MessageEmbedField{
			Value:  fmt.Sprintf("* **%s**", item.Name),
			Inline: false,
		}
	}

	return &discordgo.MessageEmbed{
		Title:  title,
		Fields: fields,
	}

}
