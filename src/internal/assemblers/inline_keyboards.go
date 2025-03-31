package assemblers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	callbackdata "github.com/jonleeyz/bball8bot/internal/data/callback-data"
)

// AssembleTrainingPollInlineKeyboard builds a basic inline keyboard for the training poll template messsage.
func AssembleTrainingPollInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Attending", callbackdata.ATTENDING),
			tgbotapi.NewInlineKeyboardButtonData("TBC, will update", callbackdata.CONFIRMING),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Not attending", callbackdata.NOT_ATTENDING),
		),
	)
}
