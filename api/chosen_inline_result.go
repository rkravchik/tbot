package api

// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ID is the unique identifier for the result that was chosen
	ID string `json:"result_id"`

	// From is the user that chose the result
	From User `json:"from"`

	// Location is a sender location, only for bots that require user location
	// Optional
	Location Location `json:"location"`

	// InlineMessageID is the identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	// Optional
	InlineMessageID string `json:"inline_message_id"`

	// Query that was used to obtain the result
	Query string `json:"query"`
}
