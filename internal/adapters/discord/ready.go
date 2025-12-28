package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Ready обработчик события готовности
func (b *Bot) Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Бот %s запущен! Префикс: '%s'\n", r.User.Username, b.prefix)
	err := s.UpdateGameStatus(0, fmt.Sprintf("Используй %shelp", b.prefix))
	if err != nil {
		log.Println(err)
	}
}
