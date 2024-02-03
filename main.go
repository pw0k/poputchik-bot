package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	//todo update config timeout for api call ??
	botToken := os.Getenv("STRANGER_TELEGRAMBOT_TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot user: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		chatID := tu.ID(message.Chat.ID)
		_, _ = bot.CopyMessage(
			tu.CopyMessage(chatID, chatID, message.MessageID),
		)
	})

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
}
