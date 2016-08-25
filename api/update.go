package api

// Update represents an incoming update.
// Only one of the optional parameters can be present in any given update.
type Update struct {
	// ID is the update‘s unique identifier.
	// Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you’re using Webhooks,
	// since it allows you to ignore repeated updates or to restore the correct update sequence,
	// should they get out of order.
	ID int32 `json:"update_id"`

	// Message is a new incoming message of any kind — text, photo, sticker, etc.
	// Optional
	Message Message `json:"message"`

	// EditedMessage is a new version of a message that is known to the bot and was edited
	// Optional
	EditedMessage Message `json:"edited_message"`

	// InlineQuery is a new incoming inline query
	// Optional
	InlineQuery InlineQuery `json:"inline_query"`

	// ChosenInlineResult is the result of an inline query that was chosen by a user and sent to their chat partner.
	// Optional
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`

	// CallbackQuery is a new incoming callback query
	// Optional
	CallbackQuery CallbackQuery `json:"callback_query"`
}
