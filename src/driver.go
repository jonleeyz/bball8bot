package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/commands"
	"github.com/jonleeyz/bball8bot/internal/handlers/callbacks"
	"github.com/jonleeyz/bball8bot/internal/handlers/messages"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

// handleUpdate parses the input Update object, and responds accordingly.
// No error is returned from this function, any error is logged but is not fatal; other updates can still be processed.
func handleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// TODO @jonlee: Update, placeholder, just to ensure that callback queries are answered.
	if update.CallbackQuery != nil {
		handler := callbacks.Init(bot)
		handler.Handle(ctx, update)
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

	messageHandler := messages.Init(bot)
	messageHandler.Handle(ctx, update)
}
