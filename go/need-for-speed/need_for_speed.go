package speed

// Car is the properties of a car.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	const initialValueBattery = 100
	const initialValueDistance = 0

	return Car{
		battery:      initialValueBattery,
		distance:     initialValueDistance,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

// Track is the properties of a track.
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	remainingAmountOfGasoline := car.battery - car.batteryDrain
	canDrive := remainingAmountOfGasoline > 0

	if !canDrive {
		return car
	}

	car.distance = car.distance + car.speed
	car.battery = remainingAmountOfGasoline

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drivableDistance := (car.battery / car.batteryDrain) * car.speed

	return drivableDistance >= track.distance
}
