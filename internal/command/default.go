package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		log.Println(err)
	}
}
