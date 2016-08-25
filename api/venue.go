package api

// Venue represents a venue
type Venue struct {
    // Location of the Venue
    Location Location `json:"location"`
    
    // Title is the	name of the venue
    Title string `json:"title"`
    
    // Address of the venue
    Address string `json:"address"`
    
    // FoursquareID is an identifier of the venue
    // Optional
    FoursquareID string `json:"foursquare_id"`
}