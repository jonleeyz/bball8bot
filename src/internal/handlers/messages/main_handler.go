package messages

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/commands"
	customerrors "github.com/jonleeyz/bball8bot/internal/custom-errors"
	"github.com/jonleeyz/bball8bot/internal/util/logging"
)

type MessageHandler struct {
	bot    *tgbotapi.BotAPI
	update *tgbotapi.Update
}

func Init(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (*MessageHandler, error) {
	if bot == nil {
		return nil, fmt.Errorf("error when creating messages handler: %s", customerrors.ERROR_MESSAGE_NIL_INPUT_BOT)
	}

	return &MessageHandler{bot: bot, update: update}, nil
}

func (h *MessageHandler) Handle(ctx context.Context) error {
	// if message is command, call command handler
	if h.update.Message.IsCommand() {
		if err := commands.HandleBotCommand(ctx, h.bot, h.update); err != nil {
			// TODO @jonlee: Tidy this log statement
			logging.Errorf("TEMP TOP level log: %v", err)
			return err
		}
		return nil
	}

	// if message is not command, echo message as reply to original message
	return h.echoMessageAsReply(ctx)
}

func (h *MessageHandler) echoMessageAsReply(ctx context.Context) error {
	newReply := tgbotapi.NewMessage(h.update.Message.Chat.ID, h.update.Message.Text)
	newReply.BaseChat.ReplyToMessageID = h.update.Message.MessageID
	if _, err := h.bot.Send(newReply); err != nil {
		logging.Errorf("error when calling Telegram Bot API to send message: %v", err)
		return err
	}
	return nil
}
