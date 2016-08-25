package api

// Video represents a video file
type Video struct {
    // ID is a unique identifier for this file
    ID string `json:"file_id"`

    // Width of the Video as defined by sender
    Width int32 `json:"width"`

    // Height of the Video as defined by sender
    Height int32 `json:"height"`

    // Duration of the video in seconds as defined by sender
    Duration int32 `json:"duration"`

    // Thumb of the Video
    // Optional
    Thumb PhotoSize `json:"thumb"`

    // MIMEType of the file as defined by sender
    // Optional
    MIMEType string `json:"mime_type"`

    // Size itself
    // Optional
    Size int32 `json:"file_size"`
}