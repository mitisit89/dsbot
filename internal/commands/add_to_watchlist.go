package commands

import (
	"context"
	"dsbot/internal/storage"
	"dsbot/internal/yt"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func AddToWathclist(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var msgformat string
	opts := i.ApplicationCommandData().Options
	optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		optsMap[opt.Name] = opt
	}
	args := make([]string, 0, len(optsMap))
	trailer, err := yt.SearchFilmTrailer(optsMap["movie"].StringValue())
	if err != nil {
		slog.Error("failed to get trailer", err)
		msgformat += fmt.Sprintf("Failed to get trailer for %s", optsMap["movie"].StringValue())
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat},
		})
		if err != nil {
			slog.Error("failed to respond", err)
		}
		return
	}
	args = append(args, optsMap["movie"].StringValue(), "https://www.youtube.com/watch?v="+trailer.Items[0].ID.VideoID)
	if err := storage.New().Add(context.Background(), i.Member.User.Username, args); err != nil {
		slog.Error("failed to add movie", err)
	}
	msgformat += fmt.Sprintf("Added %s to watchlist", optsMap["movie"].StringValue())
	slog.Info(msgformat)
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msgformat},
	})
	if err != nil {
		slog.Error("failed to respond", err)
	}

}
