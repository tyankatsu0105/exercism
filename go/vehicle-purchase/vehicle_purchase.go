package purchase

import (
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicles := []string{option1, option2}
	sort.Strings(vehicles)
	
	return vehicles[0] + " " + "is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percent := 1.0

	switch {
	case age <= 3:
		percent = 0.8
	case 5 < age && age < 10:
		percent = 0.7
	case 10 <= age:
		percent = 0.5
	}

	return originalPrice * percent
}
