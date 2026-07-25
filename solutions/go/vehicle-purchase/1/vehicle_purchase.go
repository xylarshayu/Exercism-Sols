package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	choice := option1
	if option2 < option1 {
		choice = option2
	} else {
		choice = option1
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	currPrice := originalPrice
	if age > 0.0 && age < 3.0 {
		currPrice = originalPrice * 0.8
	} else if age < 10.0 {
		currPrice = originalPrice * 0.7
	} else {
		currPrice = originalPrice * 0.5
	}
	return currPrice
}
