package handler

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"stranger-bot/internal/service"
)

type MessageHandler struct {
	bot         *telego.Bot
	userService *service.UserService
}

func NewMessageHandler(bot *telego.Bot, userService *service.UserService) *MessageHandler {
	return &MessageHandler{
		bot:         bot,
		userService: userService,
	}
}

func (h *MessageHandler) HandleMessage(message telego.Message) {
	username := message.From.Username
	h.userService.ProcessUser(username)

	// Echoing the same message back
	chatID := tu.ID(message.Chat.ID)
	_, err := h.bot.CopyMessage(
		tu.CopyMessage(chatID, chatID, message.MessageID),
	)
	if err != nil {
		fmt.Println(err)
	}
}
