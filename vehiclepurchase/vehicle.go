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
	// todo
	return 0.0
}
