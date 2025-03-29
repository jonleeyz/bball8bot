package callbacks

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/util/logging"
)

func (h *CallbackQueryHandler) handleAttendingCallback(ctx context.Context) error {
	// TODO @jonlee: Make async
	// TODO @jonlee: Error handling required?
	h.addAttendeeNameToPollMessageBody(ctx)
	h.answerAttendingCallback(ctx)
	return nil
}

func (h *CallbackQueryHandler) addAttendeeNameToPollMessageBody(ctx context.Context) error {
	// create edit message
	attendeeName := h.getAttendeeName()
	editMessage := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      h.callbackQuery.From.ID,
			MessageID:   h.callbackQuery.Message.MessageID,
			ReplyMarkup: h.callbackQuery.Message.ReplyMarkup,
		},

		// TODO @jonlee: Ideally should not append by referring to previous text; should reconstruct and update poll from stored state.
		Text:      appendAttendeeNameToAttendingSection(h.callbackQuery.Message.Text, attendeeName),
		ParseMode: tgbotapi.ModeMarkdownV2,
		// TODO @jonlee: Try entities
	}

	// send edit message
	_, err := h.bot.Send(editMessage)
	logging.Debugf("edit message attempted; edit message body: %+v", editMessage)

	if err != nil {
		// TODO @jonlee: Is this error logging necessary?
		logging.Errorf("error when editing poll message body: %v", err)
		return err
	}
	return nil
}

func (h *CallbackQueryHandler) answerAttendingCallback(ctx context.Context) error {
	// create callback answer
	callbackAnswerString := fmt.Sprintf("Poll updated: You have indicated as ATTENDING!")
	callbackAnswer := tgbotapi.NewCallback(h.callbackQuery.ID, callbackAnswerString)

	// send callback answer
	_, err := h.bot.Request(callbackAnswer)
	logging.Debugf("callback answer attempted; callback answer object: %+v", callbackAnswer)

	if err != nil {
		// TODO @jonlee: Is this error logging necessary?
		logging.Errorf("error when answering callback: %v", err)
		return err
	}
	return nil
}

// TODO @jonlee: To properly implement
func appendAttendeeNameToAttendingSection(pollMessageBody, attendeeName string) string {
	return fmt.Sprintf("%s\n%s", pollMessageBody, attendeeName)
}

// TODO @jonlee: To implement dynamically
func (h *CallbackQueryHandler) getAttendeeName() string {
	return h.callbackQuery.From.UserName
}
