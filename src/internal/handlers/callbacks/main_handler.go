package callbacks

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

type CallbackQueryHandler struct {
	bot *tgbotapi.BotAPI
}

func Init(bot *tgbotapi.BotAPI) *CallbackQueryHandler {
	return &CallbackQueryHandler{
		bot: bot,
	}
}

func (h *CallbackQueryHandler) Handle(ctx context.Context, update *tgbotapi.Update) error {
	callback := update.CallbackQuery
	callbackResponseString := fmt.Sprintf("button pressed: %s", callback.Data)

	var callbackAnswer tgbotapi.CallbackConfig
	if callbackResponseString == "button pressed: ATTENDING" {
		callbackAnswer = tgbotapi.NewCallbackWithAlert(callback.ID, callbackResponseString)
	} else {
		callbackAnswer = tgbotapi.NewCallback(callback.ID, callbackResponseString)
	}

	if _, err := h.bot.Request(callbackAnswer); err != nil {
		logging.Errorf("error when answering callback: %v", err)
		return err
	}
	return nil
}
