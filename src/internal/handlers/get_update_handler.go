package handlers

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/handlers/callbacks"
	"github.com/jonleeyz/bball8bot/internal/handlers/messages"
)

/**
 * CONTEXT: The Telegram Update object
 *
 * The Update object is a construct defined by Telegram, and structurally it could contain up to 20-odd fields.
 * In a valid Update object, exactly 1 of these fields is guaranteed to be non-nil, and the other fields are guaranteed to be nil.
 * Reference: (https://core.telegram.org/bots/api#update)
 */

// getUpdateHandler returns a new handler instance suitable to handle the input Update's content.
// Pre-condition: Exactly 1 optional field in an Update object will be non-nil.
// Updates provided by Telegram backend are guaranteed to fulfill this pre-condition.
func getUpdateHandler(ctx context.Context, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (UpdateHandler, error) {
	if isUpdateACallbackQuery(update) {
		return callbacks.Init(bot)
	}
	if isUpdateAMessage(update) {
		return messages.Init(bot)
	}
	return nil, fmt.Errorf("no suitable handler found for input update: %+v", update)
}

// isUpdateAMessage returns true if the input Update object contains a non-nil Message object, and false otherwise.
func isUpdateAMessage(update *tgbotapi.Update) bool {
	return update.Message != nil
}

// isUpdateACallbackQuery returns true if the input Update object contains a non-nil CallbackQuery object, and false otherwise.
func isUpdateACallbackQuery(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil
}

/**
 * List of Update object optional fields:
 * - message
 * - callback_query
 * ==== (unimplemented) ====
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
