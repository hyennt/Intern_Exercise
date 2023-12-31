package main

import (
	"encoding/json"
	"fmt"
	"os"

	"Exercise/models"
	"Exercise/services/offers"
	"Exercise/ultis/dataimporting"
)

const (
	outputPath = "storages/output/output.json"
	fileMode   = 0644
)

type Offers struct {
	Offers []models.Offer `json:"offers"`
}

func main() {
	// ==============START HERE==============
	var offers Offers

	// Paste param from command line
	checkInDate := os.Args[1]
	inputPath := os.Args[2]

	//Read JSON File
	offerInBytes, err := ultis.ReadJSONFile(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	err = json.Unmarshal(offerInBytes, &offers)
	if err != nil {
		fmt.Println("Error unmarshaling JSON file:", err)
	}

	// Filter Handler
	offerService := services.NewOfferService()
	filterOffer, err := offerService.FilterOffers(offers.Offers, checkInDate)
	if err != nil {
		fmt.Println(err)
	}

	// Write JSON File
	ultis.WriteJSONFile(outputPath, filterOffer, fileMode)

}
