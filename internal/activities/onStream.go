package activities

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	userStreamActivity = make(map[string]string)
	mu                 sync.Mutex
)

func OnStream(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	trackedUser := os.Getenv("STREAMER")
	streamNotificationChannel := os.Getenv("STREAM_NOTIFICATION_CHANNEL")
	if p.User.ID != trackedUser {
		return
	}
	fmt.Println("Stream started")
	mu.Lock()
	defer mu.Unlock()

	for _, activity := range p.Presence.Activities {
		if activity.Type == discordgo.ActivityTypeStreaming {
			currentStream := fmt.Sprintf("%s - %s", activity.Details, activity.URL)
			if lastStream, exists := userStreamActivity[p.User.ID]; exists && lastStream == currentStream {
				return
			}
			userStreamActivity[p.User.ID] = currentStream
			message := fmt.Sprintf("📢 %s начал стрим: **%s**\n🔗 %s", p.User.Username, activity.Details, activity.URL)
			log.Println(message)

			_, err := s.ChannelMessageSend(streamNotificationChannel, message)
			if err != nil {
				log.Println("Ошибка при отправке сообщения:", err)
			}
		}
	}
}
