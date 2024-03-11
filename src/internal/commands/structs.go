package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var handlerMap map[string]CommandHandler = map[string]CommandHandler{}

// CommandHandlerImpl contains the common components required by a CommandHandler interface implementation.
type CommandHandlerImpl struct {
	command string
	bot     *tgbotapi.BotAPI
	update  *tgbotapi.Update
}
