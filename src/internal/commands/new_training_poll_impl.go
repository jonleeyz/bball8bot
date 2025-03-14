package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jonleeyz/bbball8bot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle initiates the handling flow for the "/newtrainingpoll" command.
func (h *NewTrainingPollCommandHandlerImpl) Handle(ctx context.Context) error {
	trainingPollMessageContent, err := buildTrainingPollMessageContent(ctx, h.update)
	if err != nil {
		return err
	}

	logging.Debugf("Update.Message.ChatID: %d", h.update.Message.Chat.ID)

	trainingPollMessageResponse := tgbotapi.NewMessage(h.update.Message.Chat.ID, trainingPollMessageContent)
	trainingPollMessageResponse.ParseMode = "MarkdownV2"
	trainingPollMessageResponse.ReplyMarkup = buildInlineKeyboard()

	if _, err := h.bot.Send(trainingPollMessageResponse); err != nil {
		logging.Errorf(
			"error when calling Telegram Bot API to send /newtrainingpoll response.\n MessageConfig object: %+v",
			trainingPollMessageResponse,
		)
		return err
	}
	return nil
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

// buildInlineKeyboard builds a basic inline keyboard for the training poll template messsage.
func buildInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("test1", ""),
			tgbotapi.NewInlineKeyboardButtonData("test2", ""),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("test3", ""),
			tgbotapi.NewInlineKeyboardButtonData("test4", ""),
		),
	)
}

type trainingPollContent struct {
	day      string
	date     string
	time     string
	location string
}

const TRAINING_POLL_TEMPLATE = "*Regular practice\n%s, %s, \n%s\n%s*\n==========\n\n\n*Attending:*\n\n\n*Not attending:*\n\n\n*Checking availability:*\n\n\n*Yet to respond:*\n\n\n"
