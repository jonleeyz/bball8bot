package callbacks

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	customerrors "github.com/jonleeyz/bball8bot/internal/custom-errors"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

type CallbackQueryHandler struct {
	bot           *tgbotapi.BotAPI
	callbackQuery *tgbotapi.CallbackQuery
	update        *tgbotapi.Update
}

func Init(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (*CallbackQueryHandler, error) {
	if bot == nil {
		return nil, fmt.Errorf("error when creating callback query handler: %s", customerrors.ERROR_MESSAGE_NIL_INPUT_BOT)
	}
	if update == nil {
		return nil, fmt.Errorf("error when creating callback query handler: %s", customerrors.ERROR_MESSAGE_NIL_INPUT_UPDATE)
	}
	if update.CallbackQuery == nil {
		return nil, fmt.Errorf("error when creating callback query handler: %s || update: %+v",
			"input update has nil callback",
			*update)
	}

	return &CallbackQueryHandler{
		bot:           bot,
		callbackQuery: update.CallbackQuery,
		update:        update,
	}, nil
}

func (h *CallbackQueryHandler) Handle(ctx context.Context) error {
	callbackData := h.callbackQuery.Data
	if callbackData == CALLBACK_DATA_ATTENDING {
		return h.handleAttendingCallback(ctx)
	}

	callback := h.update.CallbackQuery
	callbackResponseString := fmt.Sprintf("button pressed: %s", callback.Data)

	callbackAnswer := tgbotapi.NewCallback(callback.ID, callbackResponseString)

	if _, err := h.bot.Request(callbackAnswer); err != nil {
		logging.Errorf("error when answering callback: %v", err)
		return err
	}
	return nil
}
