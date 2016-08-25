package api

// UserProfilePhotos represents a user's profile pictures
type UserProfilePhotos struct {
    // TotalCount is a total number of profile pictures the target user has
    TotalCount int32 `json:"total_count"`
    
    // Photos is requested profile pictures (in up to 4 sizes each)
    Photos [][]PhotoSize `json:"photos"`
}