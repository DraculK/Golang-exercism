package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if(speed == 0){
        return 0.0
    }else if(speed>=1 && speed <=4){
    	return 1
    }else if(speed>=5 && speed <=8){
    	return 0.9
    }else{
    	return 0.77
    }
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
    speedFloat := float64(speed)
	if(speed == 0){
        return (speedFloat * 221)
    }else if(speed>=1 && speed <=4){
    	return (speedFloat * 221)
    }else if(speed>=5 && speed <=8){
    	return (speedFloat * 221 * 0.9)
    }else{
    	return (speedFloat * 221 * 0.77)
    }
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
    perMinute := float64(221)/float64(60)
    speedFloat := float64(speed)
	if(speed == 0){
        return 0
    }else if(speed>=1 && speed <=4){
    	total := perMinute * speedFloat
    	totalInt := int(total)
        return totalInt
    }else if(speed>=5 && speed <=8){
    	total := perMinute * speedFloat *0.9
    	totalInt := int(total)
        return totalInt
    }else{
    	total := perMinute * speedFloat * 0.77
    	totalInt := int(total)
        return totalInt
    }
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if(speed == 0){
        return 0.0
    }else if(speed>=1 && speed <=4){
    	total := float64(speed)*float64(221)
        if(total >= limit){
            return limit
        }else{
        	return total
        }
    }else if(speed>=5 && speed <=8){
    	total := float64(speed)*float64(221)
        if(total >= limit){
            return limit
        }else{
        	return total
        }
    }else{
    	total := float64(speed)*float64(221)
        if(total >= limit){
            return limit
        }else{
        	return total
        }
    }
}