package handlers

import (
	"context"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/handlers/callbacks"
	"github.com/jonleeyz/bball8bot/internal/handlers/messages"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

// HandleUpdate parses the input Update object, and responds accordingly.
// No error is returned from this function, any error is logged but is not fatal; other updates can still be processed.
// TODO @jonlee: Change signature to return error; not returning because of Telegram's requirement does not allow the function to be general purpose
func HandleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	h, err := getUpdateHandler(ctx, bot, update)
	if err != nil {
		return
	}
	h.Handle(ctx, update)
}

// Pre-condition: Exactly 1 optional field in an Update object will be non-nil.
// Reference: https://core.telegram.org/bots/api#update
// getUpdateHandler returns a new handler appropriate to handle the input Update's content.
func getUpdateHandler(ctx context.Context, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (UpdateHandler, error) {
	if isUpdateACallbackQuery(ctx, update) {
		return callbacks.Init(bot)
	}
	if isUpdateAMessage(ctx, update) {
		return messages.Init(bot)
	}

	logging.Errorf("no appropriate update handler found; update: %+v", *update)
	return nil, errors.New("no matching handler found")
}

type UpdateHandler interface {
	Handle(ctx context.Context, update *tgbotapi.Update) error
}

/**
 * Optional fields:
 * - message
 * - edited_message
 * - channel_post
 * - edited_channel_post
 * - busines_connection
 * - business_message
 * - edited_business_message
 * - deleted_business_messages
 * - message_reaction
 * - message_reaction_count
 * - inline_query
 * - chosen_inline_result
 * - callback_query
 * - shipping_query
 * - pre_checkout_query
 * - purchased_paid_media
 * - poll
 * - poll_answer
 * - my_chat_member
 * - chat_member
 * - chat_join_request
 * - chat_boost
 * - removed_chat_boost
 */
