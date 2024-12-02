package dsbot

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
	//	{
	//	        Name: "want-to-sleep",
	//	        Description: "special command for streamer",
	//	    }
}
