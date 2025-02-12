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
	dsSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := dsbot.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	dsSession.AddHandler(activities.OnStream)
}

func main() {
	logger := dsbot.SetUpLogger()
	dsSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		slog.Info("Logged in as: %s#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := dsSession.Open()
	if err != nil {
		logger.Error("Cannot open the session: %v", err)
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
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := dsSession.ApplicationCommandDelete(dsSession.State.User.ID, *GuildID, v.ID)
			if err != nil {
				logger.Error("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	logger.Info("Gracefully shutting down.")
}
