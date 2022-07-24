package vehiclepurchase

import "strings"

func needsLicense(kind string) bool {
	if strings.ToLower(kind) == "car" || strings.ToLower(kind) == "truck" {
		return true
	} else {
		return false
	}
}

func chooseVehicle(option1, option2 string) string {
	// lexicographic sorting is the default sorting for strings
	if option1 <= option2 {
		return option1
	} else {
		return option2
	}
}

func calculateResellPrice(price uint16, age uint8) float64 {
	var estimatedPrice float64

	if age < 3 {
		estimatedPrice = float64(price) * 0.8
	} else if age >= 3 && age < 10 {
		estimatedPrice = float64(price) * 0.7
	} else { // age >= 10
		estimatedPrice = float64(price) * 0.5
	}
	return estimatedPrice
}
