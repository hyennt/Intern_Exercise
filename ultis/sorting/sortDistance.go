package ultis

import (
	"math/rand"

	"Exercise/models"
)

func FindMinDistanceMerchants(object []models.Merchant) *models.Merchant {
	minDistance := object[0].Distance
	minIndex := 0
	for i, merchant := range object {
		if merchant.Distance < minDistance {
			minDistance = merchant.Distance
			minIndex = i
		}
	}
	return &object[minIndex]
}

func QuickSortOffers(object []models.Offer) []models.Offer {
	if len(object) < 2 {
		return object
	}

	left := 0
	right := len(object) - 1

	pivot := rand.Int() % len(object)

	object[pivot], object[right] = object[right], object[pivot]

	for i, _ := range object {
		if object[i].Merchants[0].Distance < object[right].Merchants[0].Distance {
			object[left], object[i] = object[i], object[left]
			left++
		}
	}

	object[left], object[right] = object[right], object[left]

	QuickSortOffers(object[:left])
	QuickSortOffers(object[left+1:])

	return object
}
