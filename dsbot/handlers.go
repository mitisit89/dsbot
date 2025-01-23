package dsbot

import (
	"context"
	storage "dsbot/dsbot/storage/pg"
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
		// margs := make([]interface{}, 0, len(opts))
		if option, ok := optsMap["movie"]; ok {
			c := storage.New()
			if err := c.Add(context.TODO(), option.StringValue()); err != nil {
				slog.Error("failed to add movie", err)
			}

			// Option values must be type asserted from interface{}.
			// Discordgo provides utility functions to make this simple.
			msgformat += fmt.Sprintf("Added %s to watchlist", option.StringValue())
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msgformat},
		})

	},
	"show-watchlist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		c := storage.New()
		watchlist, err := c.GetAll(context.TODO())
		if err != nil {
			slog.Error("failed to get watchlist ", err)
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					"Your watchlist:\n%s",
					toMdList(watchlist.Names),
				),
			},
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
	// TODO:
	// "want-to-sleep": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	//            // Ignore type for now, they will be discussed in "responses"
	//            Type: discordgo.InteractionResponseChannelMessageWithSource,
	//            Data: &discordgo.InteractionResponseData{
	//                Content: "сегодня 😴",
	//            },
	//        })
	//    },
	// }

}
