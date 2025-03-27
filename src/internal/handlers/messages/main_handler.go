package messages

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/commands"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

type MessageHandler struct {
	bot *tgbotapi.BotAPI
}

func Init(bot *tgbotapi.BotAPI) *MessageHandler {
	return &MessageHandler{
		bot: bot,
	}
}

func (h *MessageHandler) Handle(ctx context.Context, update *tgbotapi.Update) error {
	// if message is command, call command handler
	if update.Message.IsCommand() {
		if err := commands.HandleBotCommand(ctx, h.bot, update); err != nil {
			// TODO @jonlee: Tidy this log statement
			logging.Errorf("TEMP TOP level log: %v", err)
			return err
		}
		return nil
	}

	// if message is not command, echo message as reply to original message
	return h.echoMessageAsReply(ctx, update)
}

func (h *MessageHandler) echoMessageAsReply(ctx context.Context, update *tgbotapi.Update) error {
	newReply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	newReply.BaseChat.ReplyToMessageID = update.Message.MessageID
	if _, err := h.bot.Send(newReply); err != nil {
		logging.Errorf("error when calling Telegram Bot API to send message: %v", err)
		return err
	}
	return nil
}
