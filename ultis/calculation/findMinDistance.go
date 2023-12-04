package ultis

import (
	"Exercise/models"
	"fmt"
	"sort"
)

type MinDistance struct {
	OfferID     int
	MinDistance float64
}

func FindOfferByID(offerID int, offers []models.Offer) *models.Offer {
	for _, offer := range offers {
		if offer.OfferID == offerID {
			return &offer
		}
	}
	return nil
}

func FindMinDistance(distances []float64) float64 {
	if len(distances) == 0 {
		return 0
	}
	minDistance := distances[0]
	for _, distance := range distances {
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func FindMinDistanceFromMerchants(merchants []models.Merchant) *models.Merchant {
	if len(merchants) == 0 {
		return nil
	}
	minDistance := merchants[0].Distance
	minMerchants := &merchants[0]

	for _, merchant := range merchants {
		if merchant.Distance < minDistance {
			minDistance = merchant.Distance
			minMerchants = &merchant
		}
	}
	return minMerchants

}

func FindTop2Merchant(merchants []models.Merchant) models.Merchant {
	minDistances := make([]float64, len(merchants))
	for i, merchant := range merchants {
		minDistances[i] = merchant.Distance
	}
	minDistance := FindMinDistance(minDistances)
	fmt.Println("Min Distance: ", minDistance)
	return models.Merchant{}
}

func FindTop2OffersWithMinDistance(offers []models.Offer) []models.Offer {
	minDistances := make([]MinDistance, 0)

	for _, offer := range offers {
		distances := make([]float64, len(offer.Merchants))

		for i, merchant := range offer.Merchants {
			distances[i] = merchant.Distance
			fmt.Print("Distance: ", distances[i], " ")

		}
		minDistance := FindMinDistance(distances)
		minDistances = append(minDistances, MinDistance{OfferID: offer.OfferID, MinDistance: minDistance})
	}

	sort.Slice(minDistances, func(i, j int) bool {
		return minDistances[i].MinDistance < minDistances[j].MinDistance
	})

	top2Offers := make([]models.Offer, 0)

	for _, minDistance := range minDistances {
		offer := FindOfferByID(minDistance.OfferID, offers)
		if offer != nil {
			top2Offers = append(top2Offers, *offer)
		}
	}

	return top2Offers
}
