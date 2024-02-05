package command

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// StartCommand handles the "/start" command.
type StartCommand struct{}

// Handle sends a welcome message and shows available options.
func (c *StartCommand) Handle(bot *telego.Bot, message *telego.Message) {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("Callback data button 1"). // Column 1
										WithCallbackData("callback_1"),
			tu.InlineKeyboardButton("Callback data button 2"). // Column 2
										WithCallbackData("callback_2"),
		),
		tu.InlineKeyboardRow( // Row 2
			tu.InlineKeyboardButton("URL button").WithURL("https://example.com"), // Column 1
		),
	)

	msgToSend := tu.Message(
		tu.ID(message.Chat.ID),
		"Welcome to the bot! Choose an option:",
	).WithReplyMarkup(inlineKeyboard)

	_, err := bot.SendMessage(msgToSend)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
