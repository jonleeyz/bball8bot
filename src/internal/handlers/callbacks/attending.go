package callbacks

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/assemblers"
	"github.com/jonleeyz/bball8bot/internal/util/logging"
)

func (h *CallbackQueryHandler) handleAttendingCallback(ctx context.Context) error {
	err := h.addAttendeeNameToPollMessageBody(ctx)
	logging.ErrorIfNonNil(err)

	err = h.answerAttendingCallback(ctx)
	logging.ErrorIfNonNil(err)

	// TODO @jonlee: Make async
	// go logging.ErrorIfNonNil(h.addAttendeeNameToPollMessageBody(ctx))
	// go logging.ErrorIfNonNil(h.answerAttendingCallback(ctx))
	return err
}

func (h *CallbackQueryHandler) addAttendeeNameToPollMessageBody(ctx context.Context) error {
	editMessage, err := assemblers.GetEditMessageConfigForAttendingCallback(ctx, h.callbackQuery)
	if err != nil {
		return err
	}

	_, err = h.bot.Send(editMessage)
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
