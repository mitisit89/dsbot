package commands

import (
	"context"
	"dsbot/internal/storage"
	"dsbot/internal/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func ShowWatchlist(s *discordgo.Session, i *discordgo.InteractionCreate) {
	c := storage.New()
	ctx := context.Background()
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		slog.Error("failed to defer response", "error", err)
		return
	}
	movies, err := c.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get watchlist", "error", err)
		_, _ = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Ошибка получения списка фильмов.",
		})
		return
	}
	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{utils.GenerateEmbed(movies, "Watchlist")},
	})
	if err != nil {
		slog.Error("failed to send follow-up message", "error", err)
	}
}
