package command

import (
	"botAdminDiscord/internal/adapters/discord"
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandlePing(b *discord.Bot, s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong! 🏓")
	if err != nil {
		log.Println(err)
	}
}
