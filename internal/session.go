package internal

import (
	"flag"
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
)

var GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

type Session struct {
	Session *discordgo.Session
	Logger  *slog.Logger
}

func New() (*Session, error) {
	logger := SetUpLogger()

	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		logger.Error("Invalid bot parameters: %v", err)
	}
	s.Identify.Intents = discordgo.IntentsGuildPresences
	bot := &Session{
		Session: s,
		Logger:  logger,
	}
	bot.registerHandlers()
	bot.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		bot.Logger.Info("Logged in as", "username", s.State.User.Username, "discriminator", s.State.User.Discriminator)
	})
	return bot, nil

}

func (s *Session) registerHandlers() {
	s.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			go h(s, i)
		}
	})
}
func (s *Session) Start() error {
	if err := s.Session.Open(); err != nil {
		s.Logger.Error("Cannot open the session", "error", err)
		return err
	}
	s.Logger.Info("Bot session started")
	return nil
}

func (s *Session) Stop() {
	s.Logger.Info("Shutting down bot session...")
	s.Session.Close()
}
func (s *Session) AddCommands() []*discordgo.ApplicationCommand {
	s.Logger.Info("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	for i, v := range Commands {
		cmd, err := s.Session.ApplicationCommandCreate(s.Session.State.User.ID, *GuildID, v)
		if err != nil {
			s.Logger.Error("Cannot create command", "name", v.Name, "error", err)
			continue
		}
		registeredCommands[i] = cmd
	}
	return registeredCommands
}

func (s *Session) RemoveCommands(commands []*discordgo.ApplicationCommand) {
	s.Logger.Info("Removing commands...")
	for _, v := range commands {
		err := s.Session.ApplicationCommandDelete(s.Session.State.User.ID, *GuildID, v.ID)
		if err != nil {
			s.Logger.Error("Cannot delete command", "name", v.Name, "error", err)
		}
	}
}
