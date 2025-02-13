package internal

import (
	"context"
	"dsbot/internal/commands"
	"dsbot/internal/storage"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"add-to-watchlist": commands.AddToWathclist,
	"show-watchlist":   commands.ShowWatchlist,
	"watched":          commands.Watched,
	// "game-list": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//
	// 	c := storage.New()
	// 	gameList, err := c.GameList(context.Background())
	// 	if err != nil {
	// 		slog.Error("failed to get game list ", err)
	// 	}
	// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	// 		// Ignore type for now, they will be discussed in "responses"
	// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 		Data: &discordgo.InteractionResponseData{
	// 			Embeds: []*discordgo.MessageEmbed{GenerateEmbed(gameList.Names, "Game List")},
	// 		},
	// 	})
	// },
	"add-to-game-list": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var msgformat string
		opts := i.ApplicationCommandData().Options
		optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
		for _, opt := range opts {
			optsMap[opt.Name] = opt
		}

		if option, ok := optsMap["game"]; ok {
			c := storage.New()
			if err := c.AddGame(context.TODO(), i.Member.User.Username, option.StringValue()); err != nil {
				slog.Error("failed to add game to list", err)
			}
			msgformat = fmt.Sprintf("> You add to game list: %s\n", option.StringValue())
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{

			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat,
			},
		})
	},
}
