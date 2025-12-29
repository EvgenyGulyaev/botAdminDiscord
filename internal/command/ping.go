package command

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	sendMessage(s, m, "Pong! 🏓")
}
