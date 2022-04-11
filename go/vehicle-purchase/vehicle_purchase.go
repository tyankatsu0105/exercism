package purchase

import (
	"regexp"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var expectKind = regexp.MustCompile("car|truck")

	return expectKind.MatchString(kind)
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicles := []string{option1, option2}
	sort.Strings(vehicles)
	return vehicles[0] + " " + "is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age <= 3:
		return originalPrice * 0.8
	case 5 < age && age < 10:
		return originalPrice * 0.7
	case 10 <= age:
		return originalPrice * 0.5
	}

	return originalPrice
}
