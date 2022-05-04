package cards

// IsValidIndex checks index is valid or not.
func IsValidIndex(slice []int, index int) bool {
	return len(slice) > index && index >= 0
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}

	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if !IsValidIndex(slice, index) {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if !IsValidIndex(slice, index) {
		slice = append(slice, value)
		return slice
	}

	slice[index] = value

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	slice = append(value, slice...)

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if !IsValidIndex(slice, index) {
		return slice
	}

	slice = append(slice[:index], slice[index+1:]...)

	return slice
}
