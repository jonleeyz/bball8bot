package commands

import (
	"context"
	"fmt"

	"github.com/jonleeyz/bbball8bot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle initiates the handling flow for the "/newtrainingpoll" command.
func (h *NewTrainingPollCommandHandlerImpl) Handle(ctx context.Context) error {
	trainingPollMessageContent, err := buildTrainingPollMessageContent(ctx, h.update)
	if err != nil {
		return err
	}

	trainingPollMessageResponse := tgbotapi.NewMessage(h.update.Message.Chat.ID, trainingPollMessageContent)

	if _, err := h.bot.Send(trainingPollMessageResponse); err != nil {
		err = fmt.Errorf("error when calling Telegram Bot API to send message: %v", err)

		logging.Printf(err.Error())
		return err
	}
	return nil
}

// buildTrainingPollMessageContent builds the content string for a training poll message.
func buildTrainingPollMessageContent(ctx context.Context, update *tgbotapi.Update) (string, error) {
	var baseContent string = TRAINING_POLL_TEMPLATE
	return baseContent, nil
}

const TRAINING_POLL_TEMPLATE = "Training: %s, %s @ %s\n---\n\n\nAttending:\n\n\nNot attending:\n\n\nChecking availability:\n\n\nYet to respond:\n\n\n"
