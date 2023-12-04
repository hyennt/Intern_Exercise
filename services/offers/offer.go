package services

import (
	"Exercise/models"
	ultis "Exercise/ultis/calculation"
	"Exercise/ultis/validation"
	"sort"
)

type OffersOutput struct {
	Offers []models.Offer `json:"offers"`
}

const CheckInDate = "2019-12-25"

func FilterOffers(offers []models.Offer) OffersOutput {
	// Validate Checkin Date && Validate Category
	var validOffers []models.Offer
	var top2Offer []models.Offer

	for _, offer := range offers {
		if validation.CheckValidDate(CheckInDate, offer.ValidDate) && !validation.IsHotelCategory(offer.Category) {

			// Find minimum distance of MERCHANTS

			minMerchant := ultis.FindMinDistanceFromMerchants(offer.Merchants)
			offer.Merchants = []models.Merchant{*minMerchant}

			validOffers = append(validOffers, offer)

			sort.Slice(validOffers, func(i, j int) bool {
				return validOffers[i].Merchants[0].Distance < validOffers[j].Merchants[0].Distance
			})
			
			// Sort distance of OFFERS
			//validOffers = ultis.FindTop2OffersWithMinDistance(validOffers)

		}
	}
	top2Offer = append(validOffers[:2])
	//return top2Offer
	return OffersOutput{Offers: top2Offer}
}
