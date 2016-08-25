package api

// Audio represents an audio file to be treated as music by the Telegram clients
type Audio struct {
    // ID is a unique identifier for this file
    ID string `json:"file_id"`
    
    // Duration of the audio in seconds as defined by sender
    Duration int32 `json:"duration"`
    
    // Performer of the audio as defined by sender or by audio tags
    // Optional
    Performer string `json:"performer"`
    
    // Title of the audio as defined by sender or by audio tags
    // Optional
    Title string `json:"title"`
    
    // MIMEType of the file as defined by sender
    // Optional
    MIMEType string `json:"mime_type"`
    
    // Size of the file
    // Optional
    Size int32 `json:"file_size"`
}