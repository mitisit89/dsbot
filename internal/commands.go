package internal

import (
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
	{Name: "game-list", Description: "List of games"},
	{Name: "add-to-game-list", Description: "Add game to list",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "game",
				Description: "Name of the game",
				Required:    true,
			},
		}},
	{Name: "announce", Description: "Make announce ", Options: []*discordgo.ApplicationCommandOption{
		{Type: discordgo.ApplicationCommandOptionString, Name: "time", Description: "Your local time:example 13:00", Required: true},
		{Type: discordgo.ApplicationCommandOptionString, Name: "month-day", Description: "Your local day-month: example 12-01", Required: true},
		{Type: discordgo.ApplicationCommandOptionString, Name: "description", Description: "Announce  description", Required: true},
		{Type: discordgo.ApplicationCommandOptionString, Name: "link", Description: "Link to traller", Required: false},
	}},
}
