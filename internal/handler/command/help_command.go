package command

import (
	"fmt"
	"github.com/mymmrac/telego"
)

// HelpCommand handles the "/help" command.
type HelpCommand struct{}

// Handle provides information on how to use the bot.
func (c *HelpCommand) Handle(bot *telego.Bot, message *telego.Message) {
	_, err := bot.SendMessage(&telego.SendMessageParams{
		ChatID: telego.ChatID{ID: message.Chat.ID},
		Text:   "Here are some instructions on how to use the bot...",
	})
	if err != nil {
		fmt.Println(err)
	}
}
