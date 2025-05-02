package activities

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type streamState struct {
	Streaming bool
	Details   string
	URL       string
}

var (
	userIsStreaming = make(map[string]bool)
	mu              sync.Mutex
	userStreamState = make(map[string]streamState)
)

func OnStream(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	var isStreaming bool
	var currentDetails, currentURL string
	trackedUser := os.Getenv("STREAMER")
	streamNotificationChannel := os.Getenv("STREAM_NOTIFICATION_CHANNEL")

	if p.User.ID != trackedUser {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for _, activity := range p.Presence.Activities {
		if activity.Type == discordgo.ActivityTypeStreaming {
			isStreaming = true
			currentDetails = activity.Details
			currentURL = activity.URL
			break
		}
	}
	prevState, hasPrev := userStreamState[p.User.ID]
	if isStreaming && (!hasPrev || !prevState.Streaming) {
		log.Println("Stream started")
		sendStreamNotification(s, streamNotificationChannel, p.User.Username, currentDetails, currentURL)
		userStreamState[p.User.ID] = streamState{Streaming: true, Details: currentDetails, URL: currentURL}
		return
	}
	if isStreaming && prevState.Streaming && (prevState.Details != currentDetails || prevState.URL != currentURL) {
		log.Println("Stream details or URL changed")
		sendStreamNotification(s, streamNotificationChannel, p.User.Username, currentDetails, currentURL)
		userStreamState[p.User.ID] = streamState{Streaming: true, Details: currentDetails, URL: currentURL}
		return
	}
	if !isStreaming && hasPrev && prevState.Streaming {
		log.Println("Stream ended")
		userStreamState[p.User.ID] = streamState{Streaming: false}
		return
	}
}

func sendStreamNotification(s *discordgo.Session, channelID, username, details, url string) {
	message := fmt.Sprintf("📢 %s начал стрим: **%s**\n🔗 %s", username, details, url)
	log.Println(message)

	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Fatalln("Error sending stream notification:", err)
	}
}
