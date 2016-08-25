package api

// Voice represents a voice note
type Voice struct {
    // ID is a unique identifier for this file
    ID string `json:"file_id"`

    // Duration of the audio in seconds as defined by sender
    Duration int32 `json:"duration"`

    // MIMEType of the file as defined by sender
    // Optional
    MIMEType string `json:"mime_type"`

    // Size itself
    // Optional
    Size int32 `json:"file_size"`
}