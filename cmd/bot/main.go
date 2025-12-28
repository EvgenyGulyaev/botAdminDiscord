package main

import (
	"botAdminDiscord/internal/adapters/discord"
	"botAdminDiscord/internal/command"
	"botAdminDiscord/internal/config"
	"log"
)

func main() {
	// Загружаем конфигурацию
	c := config.LoadConfig()

	b := discord.GetBot(c.Env["TOKEN"])
	b.RegisterCommand("ping", command.HandlePing)

	err := b.Run()
	if err != nil {
		log.Fatal(err)
	}
}
