package api

// Contact represents a phone contact
type Contact struct {
    // PhoneNumber of the contact
    PhoneNumber string `json:"phone_number"`
    
    // FirstName of the contact
    FirstName string `json:"first_name"`
    
    // LastName of the contact
    // Optional
    LastName string `json:"last_name"`
    
    // UserID in Telegram
    // Optional
    UserID int32 `json:"user_id"`
}