package commands

import (
	"dsbot/internal/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func Anonce(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	optsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opts))
	for _, opt := range opts {
		optsMap[opt.Name] = opt

	}

	link := optsMap["link"].StringValue()
	description := optsMap["description"].StringValue()
	time := optsMap["time"].StringValue()
	month_day := optsMap["month-day"].StringValue()
	dst := utils.ToUnixDiscordTimestamp(time, month_day)
	embed := utils.NewEmbed().SetTitle("Announce").SetDescription(description).SetURL(link).AddField("Time", dst).MessageEmbed
	msg, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
	slog.Info("send message", "message", msg)
	if err != nil {
		slog.Error("failed to send message", err)
	}

}
