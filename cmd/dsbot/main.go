package main

import (
	dsbot "dsbot/internal"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

// Bot parameters
var (
	RemoveCommands                 = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionViewChannel
)

func main() {
	bot, err := dsbot.New()
	if err != nil {
		panic(err)
	}
	err = bot.Start()
	if err != nil {
		panic(err)
	}
	defer bot.Stop()
	addedCommands := bot.AddCommands()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	bot.Logger.Info("Press Ctrl+C to exit")
	<-stop
	if *RemoveCommands {
		bot.RemoveCommands(addedCommands)
	}
	bot.Logger.Info("Gracefully shutting down.")
}
