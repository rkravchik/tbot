package api

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot,
// the field message will be presented. If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be presented.
type CallbackQuery struct {
	// ID is a unique identifier for this query
	ID string `json:"id"`

	// From is a Sender
	From User `json:"form"`

	// Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old
	// Optional
	Message Message `json:"message"`

	// InlineMessageID is the identifier of the message sent via the bot in inline mode, that originated the query
	// Optional
	InlineMessageID string `json:"inline_message_id"`

	// Data associated with the callback button.
	// Be aware that a bad client can send arbitrary data in this field
	Data string `json:"data"`
}
