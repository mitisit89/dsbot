package commands

import (
	"dsbot/internal/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func Anonce(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var link string
	opts := i.ApplicationCommandData().Options
	optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		optsMap[opt.Name] = opt

	}
	if optsMap["link"] == nil {
		link = ""
	} else {
		link = optsMap["link"].StringValue()
	}
	description := optsMap["description"].StringValue()
	userTime := optsMap["time"].StringValue()
	day := optsMap["day"].StringValue()
	usertimeZone := optsMap["timezone"].StringValue()
	dst, err := utils.ToUnixDiscordTimestamp(day, userTime, usertimeZone)
	if err != nil {
		slog.Error("failed to convert time", err)
	}
	embed := utils.NewEmbed().SetTitle("Announce").SetDescription(description).SetURL(link).AddField("Time", dst).MessageEmbed
	msg, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
	slog.Info("send message", "message", msg)
	if err != nil {
		slog.Error("failed to send message", err)
	}

}
