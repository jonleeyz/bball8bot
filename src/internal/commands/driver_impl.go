package commands

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleBotCommand calls the appropriate bot command handler if there is a valid bot command contained
// in the input Update object.
func HandleBotCommand(ctx context.Context, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	commandStringLiteral := update.Message.Command()
	commandHandler, ok := handlerMap[commandStringLiteral]
	if !ok {
		return fmt.Errorf("command not found in handlerMap: %s", commandStringLiteral)
	}

	if err := commandHandler.Init(bot, update); err != nil {
		return err
	}

	return commandHandler.Handle(ctx)
}

// Init intialises a command handler with references to an initialised Bot API and the input Update object.
func (c *CommandHandlerImpl) Init(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	if bot != nil {
		return fmt.Errorf("error when initialising /%s command handler: input BotAPI object is nil", c.command)
	}
	c.bot = bot

	if update == nil {
		return fmt.Errorf("error when initialising /%s command handler: input Update object is nil", c.command)
	}
	c.update = update

	return nil
}
