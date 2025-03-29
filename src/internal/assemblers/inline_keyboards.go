package assemblers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/handlers/callbacks"
)

// AssembleTrainingPollInlineKeyboard builds a basic inline keyboard for the training poll template messsage.
func AssembleTrainingPollInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Attending", callbacks.CALLBACK_DATA_ATTENDING),
			tgbotapi.NewInlineKeyboardButtonData("TBC, will update", callbacks.CALLBACK_DATA_CONFIRMING),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Not attending", callbacks.CALLBACK_DATA_NOT_ATTENDING),
		),
	)
}
