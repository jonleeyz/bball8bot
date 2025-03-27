package commands

import (
	"context"
	"fmt"

	"github.com/jonleeyz/bball8bot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle initiates the handling flow for the "/start" command.
func (c *StartCommandHandlerImpl) Handle(ctx context.Context) error {
	startMessageResponse := tgbotapi.NewMessage(c.update.Message.Chat.ID, START_COMMAND_RESPONSE_TEXT)

	if _, err := c.bot.Send(startMessageResponse); err != nil {
		err = fmt.Errorf("error when calling Telegram Bot API to send message: %v", err)

		logging.Infof(err.Error())
		return err
	}
	return nil
}

const START_COMMAND_RESPONSE_TEXT = "Hello from bball8bot, Basketball Ultimate's helper robot!"
