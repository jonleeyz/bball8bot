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

	upcomingSaturday := getUpcomingDate(time.Saturday)
	dayContent = upcomingSaturday.Weekday().String()
	y, m, d := upcomingSaturday.Date()
	dateContent = fmt.Sprintf("%s %d, %d", m, d, y)

	populatedTrainingPollTemplate := fmt.Sprintf(TRAINING_POLL_TEMPLATE, dayContent, dateContent, timeContent, locationContent)

	escapeDashPopulatedTrainingPollTemplate := strings.Replace(populatedTrainingPollTemplate, "-", "\\-", -1)
	return escapeDashPopulatedTrainingPollTemplate, nil
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

const TRAINING_POLL_TEMPLATE = "*Training: %s, %s, %s @ %s*\n---\n\n\n*Attending:*\n\n\n*Not attending:*\n\n\n*Checking availability:*\n\n\n*Yet to respond:*\n\n\n"
