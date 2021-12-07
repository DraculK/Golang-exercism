package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(t int) int {
	remaining := 40 - t
    return remaining
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	preparation := numberOfLayers*2
    return preparation
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elapsed := (numberOfLayers*2) + actualMinutesInOven
    return elapsed
}
