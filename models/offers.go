package models

type Offer struct {
	OfferID     int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    int        `json:"category"`
	Merchants   []Merchant `json:"merchants"`
	ValidDate   string     `json:"valid_to"`
}
