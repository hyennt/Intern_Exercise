package main

import (
	"encoding/json"
	"fmt"

	"Exercise/models"
	"Exercise/services/offers"
	"Exercise/ultis/dataimporting"
)

const (
	inputPath  = "storages/input/input2.json"
	outputPath = "storages/output/output.json"
	fileMode   = 0644
)

type Offers struct {
	Offers []models.Offer `json:"offers"`
}
  
func main() {
	// ==============START HERE==============

	var offers Offers
	// Read JSON File
	offerInBytes, err := ultis.ReadJSONFile(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	err = json.Unmarshal(offerInBytes, &offers)
	if err != nil {
		fmt.Println("Error unmarshaling JSON file:", err)
	}

	// Main logic filter handle
	filterOffer := services.FilterOffers(offers.Offers)
	fmt.Println("FILTERED OFFERS:", filterOffer)

	// Write JSON File
	ultis.WriteJSONFile(outputPath, filterOffer, fileMode)

	// ==============TEST FUNCTION HERE==============
	

}
