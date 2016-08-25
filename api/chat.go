package api

// Chat represents a chat
type Chat struct {
    // ID is a unique identifier for this chat
    ID int64 `json:"id"`
    
    // Type of the chat, can be either "private", "group", "supergroup" or "channel"
    Type string `json:"type"`
    
    // Title for channels and group chats
    // Optional
    Title string `json:"title"`
    
    // UserName for private chats and channels if available
    // Optional
    UserName string `json:"username"`
    
    // FirstName of the other party in a private chat
    FirstName string `json:"first_name"`
    
    // LastName of the other party in a private chat
    // Optional
    LastName string `json:"last_name"`
}