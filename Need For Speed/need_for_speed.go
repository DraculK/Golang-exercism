package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
        speed: speed,
        batteryDrain: batteryDrain,
        battery: 100,
        distance: 0,
    }
	return car
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}
// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
        distance: distance,
    }
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
    if(car.battery>=car.batteryDrain){
        car.distance+=car.speed
    	car.battery-=car.batteryDrain
    }else{
    	car.distance = 0
    }
	
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if(car.battery>=car.batteryDrain){
        for(car.distance<track.distance){
            car.distance+=car.speed
            car.battery-=car.batteryDrain
        }
    	if(car.battery<=0 || car.distance<track.distance){
            return false
        }else{
        	return true
        }
    }else{
    	return false
    }
}