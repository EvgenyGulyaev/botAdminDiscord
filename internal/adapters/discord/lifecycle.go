package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (b *Bot) Run() error {
	done, err := b.Start()
	if err != nil {
		return err
	}

	b.WaitForShutdown()
	<-done
	return nil
}

func (b *Bot) Start() (<-chan struct{}, error) {
	b.bot.AddHandler(b.Ready)
	b.bot.AddHandler(b.MessageCreate)

	err := b.bot.Open()
	if err != nil {
		return nil, err
	}

	// Запускаем горутину для ожидания сигналов
	done := make(chan struct{})
	go b.waitForShutdown(done)

	return done, nil
}

func (b *Bot) Stop() {
	close(b.stopChan)
	err := b.bot.Close()
	if err != nil {
		log.Println("Can't close bot", err)
	}
}

func (b *Bot) WaitForShutdown() {
	fmt.Println("Бот запущен. Нажмите Ctrl+C для выхода.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	b.Stop()
}

func (b *Bot) waitForShutdown(done chan<- struct{}) {
	defer close(done)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)

	select {
	case <-sc:
		fmt.Println("\nПолучен сигнал ОС. Выполняю graceful shutdown...")
	case <-b.stopChan:
		fmt.Println("Получен внутренний сигнал остановки...")
	}

	err := b.bot.Close()
	if err != nil {
		log.Println("Can't close bot", err)
	}
}
