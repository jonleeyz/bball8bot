package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	BOT_COMMAND_START             = "start"
	BOT_COMMAND_NEW_TRAINING_POLL = "newtrainingpoll"
)

var (
	// handlerMap maps bot command string literals to their appropriate handler implementation objects.
	handlerMap map[string]CommandHandler = map[string]CommandHandler{
		BOT_COMMAND_START: &StartCommandHandlerImpl{
			CommandHandlerImpl: CommandHandlerImpl{
				command: BOT_COMMAND_START,
			},
		},
		BOT_COMMAND_NEW_TRAINING_POLL: &NewTrainingPollCommandHandlerImpl{
			CommandHandlerImpl: CommandHandlerImpl{
				command: BOT_COMMAND_NEW_TRAINING_POLL,
			},
		},
	}
)

// CommandHandlerImpl contains the common components required by a CommandHandler interface implementation.
type CommandHandlerImpl struct {
	command string
	bot     *tgbotapi.BotAPI
	update  *tgbotapi.Update
}

// StartCommandHandlerImpl is a CommandHandler interface implementation that handles the "/start" command.
type StartCommandHandlerImpl struct {
	CommandHandlerImpl
}

// NewTrainingPollCommandHandlerImpl is a CommandHandler interface implementation that handles the "/start" command.
type NewTrainingPollCommandHandlerImpl struct {
	CommandHandlerImpl
}
