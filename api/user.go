package api

// User represents a Telegram user or bot
type User struct {
    // ID Unique identifier for this user or bot
    ID int32 `json:"id"`
    
    // FirstName of User or bot
    FirstName string `json:"first_name"`
    
    // LastName of User or bot
    // Optional
    LastName string `json:"last_name"`
    
    // UserName of User or bot
    // Optional
    UserName string `json:"username"`
}