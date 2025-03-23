package commands

import (
	"context"
	"dsbot/internal/storage"
	"dsbot/internal/yt"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func AddToWatchlist(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var msgFormat string
	opts := i.ApplicationCommandData().Options
	optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		optsMap[opt.Name] = opt
	}
	movieName := optsMap["movie"].StringValue()
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		slog.Error("failed to defer response", "error", err)
		return
	}
	trailer, err := yt.SearchFilmTrailer(movieName)
	if err != nil {
		slog.Error("failed to get trailer", "error", err)
		msgFormat = fmt.Sprintf("Failed to get trailer for %s", movieName)
		_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: msgFormat,
		})
		if err != nil {
			slog.Error("failed to send follow-up message", "error", err)
		}
		return
	}
	args := []string{movieName, "https://www.youtube.com/watch?v=" + trailer.Items[0].ID.VideoID}
	if err := storage.New().Add(context.Background(), i.Member.User.Username, args); err != nil {
		slog.Error("failed to add movie", "error", err)
		_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Ошибка при добавлении фильма в список.",
		})
		if err != nil {
			slog.Error("failed to send follow-up message", "error", err)
		}
		return
	}
	msgFormat = fmt.Sprintf("Added %s to watchlist", movieName)
	slog.Info(msgFormat)

	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: msgFormat,
	})
	if err != nil {
		slog.Error("failed to send follow-up message", "error", err)
	}
}
