package handler

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"stranger-bot/internal/service"
)

type MessageHandler struct {
	bot            *telego.Bot
	userService    *service.UserService
	commandService *CommandService
}

func NewMessageHandler(bot *telego.Bot, userService *service.UserService, commandService *CommandService) *MessageHandler {
	return &MessageHandler{
		bot:            bot,
		userService:    userService,
		commandService: commandService,
	}
}

func (h *MessageHandler) HandleMessage(message *telego.Message) {
	if isCommand(message) {
		h.commandService.ExecuteCommand(message.Text, h.bot, message)
		return
	}
	h.userService.ProcessUser(message.From.Username)
	sendSameMsg(message, h)
}

func isCommand(message *telego.Message) bool {
	return len(message.Text) > 0 && message.Text[0] == '/'
}

func sendSameMsg(message *telego.Message, h *MessageHandler) {
	chatID := tu.ID(message.Chat.ID)
	_, err := h.bot.CopyMessage(
		tu.CopyMessage(chatID, chatID, message.MessageID),
	)
	if err != nil {
		fmt.Println(err)
	}
}
