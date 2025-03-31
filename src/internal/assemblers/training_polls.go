package assemblers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	customerrors "github.com/jonleeyz/bball8bot/internal/util/custom-errors"
)

// GetEditMessageConfigForNewAttendee outputs the edit message object that should reflect a training poll's correctly updated
// message body after an ATTENDING callback query is received.
func GetEditMessageConfigForAttendingCallback(
	ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) (*tgbotapi.EditMessageTextConfig, error) {

	if callbackQuery == nil {
		return nil, getEditForAttendingError(customerrors.ERROR_NIL_CALLBACK)
	}
	if callbackQuery.Message == nil {
		return nil, getEditForAttendingError(customerrors.ERROR_NIL_CALLBACK_MESSAGE)
	}
	if callbackQuery.From == nil {
		return nil, getEditForAttendingError(customerrors.ERROR_NIL_CALLBACK_USER)
	}

	updatedMessageBody, err := generateUpdatedMessageBodyTextReflectingNewlyAttendingUser(
		ctx, callbackQuery.Message.Text, callbackQuery.From)
	if err != nil {
		return nil, err
	}

	editMessage := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      callbackQuery.From.ID,
			MessageID:   callbackQuery.Message.MessageID,
			ReplyMarkup: callbackQuery.Message.ReplyMarkup,
		},

		Text:      updatedMessageBody,
		ParseMode: tgbotapi.ModeMarkdownV2,
		// TODO @jonlee: Try entities
	}
	return &editMessage, nil
}

// generateUpdatedMessageBodyTextReflectingNewlyAttendingUser updates the input training poll message body text to reflect that
// the input user has indicated "attending".
// TODO @jonlee: Should not update by referring to previous text; should update poll stored state and then reconstruct.
func generateUpdatedMessageBodyTextReflectingNewlyAttendingUser(
	ctx context.Context, trainingPollMessageBodyText string, user *tgbotapi.User) (string, error) {

	if user == nil {
		return "", createNewAttendeeUpdatedMessageBodyError(customerrors.ERROR_NIL_USER)
	}

	userHandle := user.String()
	return fmt.Sprintf("%s\n%s", trainingPollMessageBodyText, userHandle), nil
}

var getEditForAttendingError = customerrors.CreateErrorTemplate("error when getting edit message for attending callback: %s")
var createNewAttendeeUpdatedMessageBodyError = customerrors.CreateErrorTemplate("error when generating updated message body with new attending user: %s")
