package services

import (
	"Exercise/constants"
	"Exercise/models"
	"Exercise/ultis/calculation"
	"Exercise/ultis/validation"
)

type OffersOutput struct {
	Offers []models.Offer `json:"offers"`
}

type IOfferService interface {
	FilterOffers(offers []models.Offer, heckInDate string) (OffersOutput, error)
	// ...
	//You can implement more logic in here in order to maintain and extend easily
}

type OfferService struct {
	// Call database and log here
}

func NewOfferService() IOfferService {
	return &OfferService{}
}

// Time : 2.313
func (o *OfferService) FilterOffers(offers []models.Offer, checkInDate string) (OffersOutput, error) {
	var validOffers []models.Offer

	// Validate CheckInDate and Category
	filterValidOffers := func(offer models.Offer) bool {
		return validation.CheckValidDate(checkInDate, offer.ValidDate) && constants.CategoriesMapping["Hotel"] != offer.Category
	}

	// Sort Merchants by distance
	processMerchants := func(offer models.Offer) models.Offer {
		minMerchants := ultis.SortMerchants(offer.Merchants)
		offer.Merchants = []models.Merchant{*minMerchants}
		return offer
	}

	for _, offer := range offers {
		if filterValidOffers(offer) {
			validOffers = append(validOffers, processMerchants(offer))
		}
	}

	// Sort Offers by Distance
	validOffers = ultis.SortValidOffers(validOffers)

	if len(validOffers) > 2 {
		validOffers = validOffers[:2]
	}

	return OffersOutput{Offers: validOffers}, nil
}
