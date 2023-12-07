package services

import (
	"errors"

	"Exercise/constants"
	"Exercise/models"
	"Exercise/ultis/sorting"
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
	// Call database and log here for
}

func NewOfferService() IOfferService {
	return &OfferService{}
}

func (o *OfferService) FilterOffers(offers []models.Offer, checkInDate string) (OffersOutput, error) {
	var validOffers []models.Offer

	checkInDateTime, ok := validation.ValidDate(checkInDate)
	if !ok {
		return OffersOutput{nil}, errors.New("FORMATTED DATE MUST BE YYYY-MM-DD")
	}

	// Validate CheckInDate and Category
	filterValidOffers := func(offer models.Offer) bool {
		return validation.CheckValidDate(checkInDateTime, offer.ValidDate) && constants.CategoriesMapping["Hotel"] != offer.Category
	}

	// Process Merchants to find min distance
	processMerchants := func(offer models.Offer) models.Offer {
		minMerchants := ultis.FindMinDistanceMerchants(offer.Merchants)
		offer.Merchants = []models.Merchant{*minMerchants}
		return offer
	}
	for _, offer := range offers {
		if filterValidOffers(offer) {
			validOffers = append(validOffers, processMerchants(offer))
		}
	}

	// Sort Offers by Distance
	validOffers = ultis.QuickSortOffers(validOffers)

	if len(validOffers) > 2 {
		validOffers = validOffers[:2]
	}

	return OffersOutput{Offers: validOffers}, nil
}
