package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
	// if knightIsAwake == false {
	// 	return true
	// }

	// return false
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
	// if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true {
	// 	return true
	// }

	// return false
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
	// if archerIsAwake == false && prisonerIsAwake == true {
	// 	return true
	// }

	// return false
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (!archerIsAwake && petDogIsPresent) || (!knightIsAwake && !archerIsAwake && prisonerIsAwake)
	// if archerIsAwake == false && petDogIsPresent == true {
	// 	return true
	// }

	// if knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true {
	// 	return true
	// }

	// return false
}
