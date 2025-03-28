package commands

import (
	"context"
	"dsbot/internal/storage"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func Watched(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var msgformat string
	opts := i.ApplicationCommandData().Options
	optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		optsMap[opt.Name] = opt
	}
	if option, ok := optsMap["movie"]; ok {
		c := storage.New()

		if err := c.MarkAsWatched(context.TODO(), option.StringValue()); err != nil {
			slog.Error("failed to mark as watched", err)
		}
		msgformat = fmt.Sprintf("> You mark as watched: %s\n", option.StringValue())

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat,
			},
		})
		if err != nil {
			slog.Error("failed to respond", err)
		}
	}

}
