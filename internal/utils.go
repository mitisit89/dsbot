package internal

import (
	"dsbot/internal/storage"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// TODO:create generic function for gen embed message
func GenerateEmbed(movies []storage.Movie, title string) *discordgo.MessageEmbed {
	fields := make([]*discordgo.MessageEmbedField, len(movies))
	for i, item := range movies {
		if item.DiscordUser.Valid && item.Trailer.Valid {
			fields[i] = &discordgo.MessageEmbedField{
				Value:  fmt.Sprintf("**%s**\n> [Trailer](%s)\n> Ordered by %s", item.Name, item.Trailer.String, item.DiscordUser.String),
				Inline: false,
			}
		} else {
			fields[i] = &discordgo.MessageEmbedField{
				Value:  fmt.Sprintf("**%s**\n", item.Name),
				Inline: false,
			}
		}
	}

	return &discordgo.MessageEmbed{
		Title:  title,
		Fields: fields,
	}

}
