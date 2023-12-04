package models

type Merchant struct {
	MerchantID int     `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
}
