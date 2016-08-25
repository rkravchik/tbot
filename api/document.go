package api

// Document represents a general file (as opposed to photos, voice messages and audio files)
type Document struct {
    // ID is a unique file identifier
    ID string `json:"file_id"`
    
    // Thumb is the document thumbnail as defined by sender
    // Optional
    Thumb PhotoSize `json:"thumb"`

    // Name is original filename as defined by sender
    // Optional
    Name string `json:"file_name"`

    // MIMEType of the file as defined by sender
    // Optional
    MIMEType string `json:"mime_type"`

    // Size of the document file
    // Optional
    Size int32 `json:"file_size"`
}