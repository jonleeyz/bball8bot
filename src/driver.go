package main

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/commands"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

// handleUpdate parses the input Update object, and responds accordingly.
// No error is returned from this function, any error is logged but is not fatal; other updates can still be processed.
func handleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// TODO @jonlee: Update, placeholder, just to ensure that callback queries are answered.
	if update.CallbackQuery != nil {
		callback := update.CallbackQuery
		callbackResponseString := fmt.Sprintf("button pressed: %s", callback.Data)

		var callbackAnswer tgbotapi.CallbackConfig
		if callbackResponseString == "button pressed: ATTENDING" {
			callbackAnswer = tgbotapi.NewCallbackWithAlert(callback.ID, callbackResponseString)
		} else {
			callbackAnswer = tgbotapi.NewCallback(callback.ID, callbackResponseString)
		}

		if _, err := bot.Request(callbackAnswer); err != nil {
			logging.Errorf("error when answering callback: %v", err)
		}
		return
	}

	if update.Message == nil {
		return
	}

	// if message is command, call command handler
	if update.Message.IsCommand() {
		if err := commands.HandleBotCommand(ctx, bot, update); err != nil {
			// TODO @jonlee: Tidy this log statement
			logging.Errorf("TEMP TOP level log: %v", err)
		}
		return
	}

	// if message is not command, echo message as reply to original message
	newReply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	newReply.BaseChat.ReplyToMessageID = update.Message.MessageID
	if _, err := bot.Send(newReply); err != nil {
		logging.Errorf("error when calling Telegram Bot API to send message: %v", err)
	}
}
