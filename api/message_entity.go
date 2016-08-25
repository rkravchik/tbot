package api

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
    // Type of the entity.
    // One of mention (@username), hashtag, bot_command, url, email,
    // bold (bold text), italic (italic text), code (monowidth string),
    // pre (monowidth block), text_link (for clickable text URLs)
    Type string `json:"type"`
    
    // Offset in UTF-16 code units to the start of the entity
    Offset int32 `json:"offset"`
    
    // Length of the entity in UTF-16 code units
    Length int32 `json:"length"`
    
    // URL for “text_link” only, url that will be opened after user taps on the text
    // Optional
    URL string `json:"url"`
    
    // User that was mentioned. For “text_mention” only
    // Optional
    User User `json:"user"`
}