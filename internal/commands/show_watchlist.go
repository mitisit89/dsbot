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
	movies, err := c.GetAll(context.Background())
	if err != nil {
		slog.Error("failed to get watchlist ", err)
	}
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{utils.GenerateEmbed(movies, "Watchlist")}},
	})
	if err != nil {
		slog.Error("failed to respond", err)
	}
}
