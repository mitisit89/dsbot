package internal

import (
	"context"
	"dsbot/internal/storage"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"add-to-watchlist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var msgformat string
		opts := i.ApplicationCommandData().Options
		optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
		for _, opt := range opts {
			optsMap[opt.Name] = opt
		}
		args := make([]string, 0, len(optsMap))
		args = append(args, optsMap["movie"].StringValue(), "trailer")
		if err := storage.New().Add(context.Background(), i.Member.User.Username, args); err != nil {
			slog.Error("failed to add movie", err)
		}
		msgformat += fmt.Sprintf("Added %s to watchlist", optsMap["movie"].StringValue())
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat},
		})

	},
	"show-watchlist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		c := storage.New()
		movies, err := c.GetAll(context.Background())
		if err != nil {
			slog.Error("failed to get watchlist ", err)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{GenerateEmbed(movies, "Watchlist")}},
		})
	},
	"watched": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var msgformat string
		opts := i.ApplicationCommandData().Options
		optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
		for _, opt := range opts {
			optsMap[opt.Name] = opt
		}
		margs := make([]interface{}, 0, len(opts))
		if option, ok := optsMap["movie"]; ok {
			c := storage.New()

			if err := c.MarkAsWatched(context.TODO(), option.StringValue()); err != nil {
				slog.Error("failed to mark as watched", err)
			}
			// Option values must be type asserted from interface{}.
			// Discordgo provides utility functions to make this simple.
			margs = append(margs, option.StringValue())
			msgformat = fmt.Sprintf("> You mark as watched: %s\n", option.StringValue())
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat,
			},
		})

	},
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
