package commands

import (
	"context"
	"fmt"
	"strings"

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
	trainingPollMessageResponse.ParseMode = "MarkdownV2"

	if _, err := h.bot.Send(trainingPollMessageResponse); err != nil {
		err = fmt.Errorf("error when calling Telegram Bot API to send message: %v", err)

		logging.Printf(err.Error())
		return err
	}
	return nil
}

// buildTrainingPollMessageContent builds the content string for a training poll message.
func buildTrainingPollMessageContent(ctx context.Context, update *tgbotapi.Update) (string, error) {
	var (
		dayContent      string = "Saturday"
		dateContent     string = "Mar 16, 2024"
		timeContent     string = "0915 - 1215"
		locationContent string = "NTU"
	)
	populatedTrainingPollTemplate := fmt.Sprintf(TRAINING_POLL_TEMPLATE, dayContent, dateContent, timeContent, locationContent)

	escapeDashPopulatedTrainingPollTemplate := strings.Replace(populatedTrainingPollTemplate, "-", "\\-", -1)
	return escapeDashPopulatedTrainingPollTemplate, nil
}

const TRAINING_POLL_TEMPLATE = "*bold \\* Training: %s, %s, %s @ %s*\n---\n\n\n*bold \\*Attending:*\n\n\n*bold \\*Not attending:*\n\n\n*bold \\*Checking availability:*\n\n\n*bold \\*Yet to respond:*\n\n\n"
