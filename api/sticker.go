package api

// Sticker represents a sticker
type Sticker struct {
    // ID is a unique identifier for this file
    ID string `json:"file_id"`

    // Width of the Sticker
    Width int32 `json:"width"`

    // Height of the Sticker
    Height int32 `json:"height"`
    
    // Thumb is a sticker thumbnail in .webp or .jpg format
    // Optional
    Thumb PhotoSize `json:"thumb"`
    
    // Emoji associated with the sticker
    // Optional
    Emoji string `json:"emoji"`
    
    // Size itself
    // Optional
    Size int32 `json:"file_size"`
}