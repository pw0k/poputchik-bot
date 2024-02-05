package handler

import "github.com/mymmrac/telego"

type CommandHandler interface {
	Handle(bot *telego.Bot, message *telego.Message)
}
