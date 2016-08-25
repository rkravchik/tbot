package api

// PhotoSize represents one size of a photo or a file / sticker thumbnail
type PhotoSize struct {
    // ID is a unique identifier for this file
    ID string `json:"file_id"`
    
    // Width of the	Photo
    Width int32 `josn:"width"`
    
    // Height of the Photo
    Height int32 `json:"height"`
    
    // Size of the Photo
    // Optional
    Size int32 `json:"file_size"`
}