package discord

import (
	"botAdminDiscord/pkg/singleton"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)

type Bot struct {
	prefix   string
	bot      *discordgo.Session
	commands map[string]CommandHandler
	mu       sync.RWMutex
	stopChan chan struct{} // Канал для остановки
}

func GetBot(botToken string) *Bot {
	return singleton.GetInstance("bot", func() interface{} {
		bot, err := initBot(botToken)
		if err != nil {
			log.Fatal("Can't start bot")
		}
		return bot
	}).(*Bot)
}

func initBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		prefix:   "!",
		bot:      dg,
		commands: make(map[string]CommandHandler),
		stopChan: make(chan struct{}),
	}, nil
}
