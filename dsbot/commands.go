package dsbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "add-to-watchlist",
		Description: "Add a anime or film to  watchlist",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "movie", // only in lover case
				Description: "Name of the anime or film",
				Required:    true,
			},
		},
	},
	{
		Name:        "show-watchlist",
		Description: "Show watchlist",
	},
	{
		Name:        "watched",
		Description: "Remove a anime or film from watchlist",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "movie", // only in lover case
				Description: "Mark as watched",
				Required:    true,
			},
		},
	},
}

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
			c := connect()
			add(c, option.StringValue())
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
		c := connect()
		watchlist := getAll(c)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			// Ignore type for now, they will be discussed in "responses"
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					"Your watchlist:\n%s",
					toMdList(watchlist),
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
			c := connect()
			remove(c, option.StringValue())
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
}
