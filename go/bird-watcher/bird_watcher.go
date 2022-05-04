package birdwatcher

import "math"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, bird := range birdsPerDay {
		total += bird
	}

	return total
}

func getBirdsPerWeek(birdsPerDay []int) map[int][]int {
	birdsPerWeek := map[int][]int{}
	for index, bird := range birdsPerDay {
		weekCount := math.Ceil(float64(index+1) / 7)
		birdsPerWeek[int(weekCount)] = append(birdsPerWeek[int(weekCount)], bird)
	}

	return birdsPerWeek
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsPerWeek := getBirdsPerWeek(birdsPerDay)

	total := 0
	for _, bird := range birdsPerWeek[week] {
		total += bird
	}

	return total

}

func isOdd(index int) bool {
	return index%2 == 0
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedBirds := birdsPerDay

	for index, bird := range birdsPerDay {

		if isOdd(index) {
			fixedBirds[index] = bird + 1
		}
	}

	return fixedBirds
}
