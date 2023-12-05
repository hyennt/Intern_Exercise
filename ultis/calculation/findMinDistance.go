package ultis

import (
	"Exercise/models"
	"sort"
)

type MinDistance struct {
	OfferID     int
	MinDistance float64
}

func SortMerchants(object []models.Merchant) *models.Merchant {
	sort.Slice(object, func(i, j int) bool {
		return object[i].Distance < object[j].Distance
	})
	return &object[0]
}

func SortValidOffers(object []models.Offer) []models.Offer {
	sort.Slice(object, func(i, j int) bool {
		return object[i].Merchants[0].Distance < object[j].Merchants[0].Distance
	})
	return object
}
