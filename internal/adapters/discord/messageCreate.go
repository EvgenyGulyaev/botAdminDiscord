package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Проверяем префикс
	if !strings.HasPrefix(m.Content, b.prefix) {
		return
	}

	content := strings.TrimPrefix(m.Content, b.prefix)
	args := strings.Fields(content)
	if len(args) == 0 {
		return
	}

	commandName := strings.ToLower(args[0])
	args = args[1:]

	b.mu.RLock()
	handler, exists := b.commands[commandName]
	b.mu.RUnlock()

	if exists && handler != nil {
		handler(s, m, args)
	}
}
