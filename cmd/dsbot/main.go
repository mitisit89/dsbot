package main

import (
	dsbot "dsbot/internal"
	"dsbot/internal/activities"
	"flag"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")

	dsSession *discordgo.Session

	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionViewChannel
)

func init() {
	logger := dsbot.SetUpLogger()
	var err error
	dsSession, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		logger.Error("Invalid bot parameters: %v", err)
	}
	dsSession.Identify.Intents = discordgo.IntentsGuildPresences
	dsSession.State.MaxMessageCount = 0 // disable message cache
	dsSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := dsbot.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			go h(s, i)
		}
	})

	dsSession.AddHandler(activities.OnStream)
}

func main() {
	logger := dsbot.SetUpLogger()
	dsSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info("Logged in as: %s#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := dsSession.Open()
	if err != nil {
		logger.Error("Cannot open the session: %v", err)
		return
	}

	logger.Info("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(dsbot.Commands))
	for i, v := range dsbot.Commands {
		cmd, err := dsSession.ApplicationCommandCreate(dsSession.State.User.ID, *GuildID, v)
		if err != nil {
			slog.Error("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer dsSession.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logger.Info("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		logger.Info("Removing commands...")
		for _, v := range registeredCommands {
			err := dsSession.ApplicationCommandDelete(dsSession.State.User.ID, *GuildID, v.ID)
			if err != nil {
				logger.Error("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	logger.Info("Gracefully shutting down.")
}
