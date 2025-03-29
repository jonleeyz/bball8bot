package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jonleeyz/bball8bot/internal/assemblers"
	"github.com/jonleeyz/bball8bot/internal/util/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle initiates the handling flow for the "/newtrainingpoll" command.
func (h *NewTrainingPollCommandHandlerImpl) Handle(ctx context.Context) error {
	// trainingPollMessageContent, err := buildTrainingPollMessageContent(ctx, h.update)
	trainingPollMessageContent, err := provideSimpleTestTrainingPollMessageContent(ctx)
	if err != nil {
		return err
	}

	logging.Debugf("Update.Message.ChatID: %d", h.update.Message.Chat.ID)

	trainingPollMessageResponse := tgbotapi.NewMessage(h.update.Message.Chat.ID, trainingPollMessageContent)
	trainingPollMessageResponse.ParseMode = "MarkdownV2"
	trainingPollMessageResponse.ReplyMarkup = assemblers.AssembleTrainingPollInlineKeyboard()

	if _, err := h.bot.Send(trainingPollMessageResponse); err != nil {
		logging.Errorf(
			"error when calling Telegram Bot API to send /newtrainingpoll response.\n MessageConfig object: %+v",
			trainingPollMessageResponse,
		)
		return err
	}
	return nil
}

// TODO @jonlee: Delete once tests are complete
// provideSimpleTestTrainingPollMessageContent builds a simple content string for testing purposes.
func provideSimpleTestTrainingPollMessageContent(ctx context.Context) (string, error) {
	messageContent := "hallo chat from Maki15Pro"
	return messageContent, nil
}

// buildTrainingPollMessageContent builds the content string for a training poll message.
func buildTrainingPollMessageContent(ctx context.Context, update *tgbotapi.Update) (string, error) {
	content := generateTrainingPollContent(time.Saturday)
	trainingPollMessageContent := fmt.Sprintf(TRAINING_POLL_TEMPLATE, content.day, content.date, content.time, content.location)
	return addEscapeTokens(trainingPollMessageContent), nil
}

// generateTrainingPollContent returns a trainingPollContent object, complete with generated content.
// For now, time and location are hardcoded.
func generateTrainingPollContent(targetWeekday time.Weekday) trainingPollContent {
	upcomingDateObject := getUpcomingDate(targetWeekday)
	y, m, d := upcomingDateObject.Date()

	return trainingPollContent{
		day:      upcomingDateObject.Weekday().String(),
		date:     fmt.Sprintf("%s %d, %d", m, d, y),
		time:     "0915 - 1215",      // TODO: Update, hardcoded for now
		location: "NTU Upper Fields", // TODO: Update, hardcoded for now
	}
}

// getUpcomingDate returns the date of the next upcoming specified weekday.
func getUpcomingDate(targetWeekday time.Weekday) time.Time {
	currentDateTime := time.Now()
	weekdayDiff := targetWeekday - currentDateTime.Weekday()
	if weekdayDiff <= 0 {
		weekdayDiff += 7
	}

	return currentDateTime.AddDate(0, 0, int(weekdayDiff))
}

// TODO: Consider using tgbotapi.EscapeText()
// addEscapeTokens adds "//" characters so the input training poll message content string can be parsed correctly
// by Telegram's Bot API.
func addEscapeTokens(trainingPollMessageContent string) string {
	trainingPollMessageContent = strings.Replace(trainingPollMessageContent, "-", "\\-", -1)
	trainingPollMessageContent = strings.Replace(trainingPollMessageContent, "=", "\\=", -1)
	return trainingPollMessageContent
}

type trainingPollContent struct {
	day      string
	date     string
	time     string
	location string
}

const TRAINING_POLL_TEMPLATE = "*Regular practice\n%s, %s, \n%s\n%s*\n==========\n\n\n*Attending:*\n\n\n*Not attending:*\n\n\n*Checking availability:*\n\n\n*Yet to respond:*\n\n\n"
const CALLBACK_QUERY_BUTTON_PRESSED = "CALLBACK_QUERY_BUTTON_PRESSED"
