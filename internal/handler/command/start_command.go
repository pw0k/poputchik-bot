package command

import (
	"fmt"
	"github.com/mymmrac/telego"
)

// StartCommand handles the "/start" command.
type StartCommand struct{}

// Handle sends a welcome message and shows available options.
func (c *StartCommand) Handle(bot *telego.Bot, message *telego.Message) {
	fmt.Println("start command handled")
}

//	replyMarkup := tu.InlineKeyboardMarkup(
//		tu.NewInlineKeyboardRow(
//			tu.NewInlineKeyboardButtonData("Help", "/help"),
//		),
//	)
//
//	_, err := bot.SendMessage(telego.SendMessageParams{
//		ChatID:      message.Chat.ID,
//		Text:        "Welcome to the bot! Here is what you can do:",
//		ReplyMarkup: &replyMarkup,
//	})
//	if err != nil {
//		// Log the error or handle it
//	}
//}
