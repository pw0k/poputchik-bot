package handler

import (
	"github.com/mymmrac/telego"
	"stranger-bot/internal/handler/command"
)

type CommandService struct {
	commands map[string]CommandHandler
}

func NewCommandService() *CommandService {
	cs := &CommandService{
		commands: make(map[string]CommandHandler),
	}
	cs.initCommands()
	return cs
}

func (cs *CommandService) initCommands() {
	cs.commands["/start"] = &command.StartCommand{}
	cs.commands["/help"] = &command.HelpCommand{}
}

func (cs *CommandService) ExecuteCommand(command string, bot *telego.Bot, message *telego.Message) {
	if cmdHandler, exists := cs.commands[command]; exists {
		cmdHandler.Handle(bot, message)
	}
}
