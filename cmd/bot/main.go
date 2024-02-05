package main

import (
	"fmt"
	"os"
	"stranger-bot/internal/config"
	"stranger-bot/internal/handler"
	"stranger-bot/internal/service"

	"github.com/mymmrac/telego"
)

func main() {
	cfg := config.LoadConfig()

	bot, err := telego.NewBot(cfg.BotToken, telego.WithDefaultLogger(false, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer bot.StopLongPolling()

	userService := service.NewUserService()
	commandService := handler.NewCommandService()
	messageHandler := handler.NewMessageHandler(bot, userService, commandService)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	for update := range updates {
		if update.Message != nil {
			messageHandler.HandleMessage(update.Message)
		}
	}
}
