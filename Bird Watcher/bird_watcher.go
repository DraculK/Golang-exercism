package birdwatcher

 TotalBirdCount return the total bird count by summing
 the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdsCount = 0
    for i=0; ilen(birdsPerDay);i++{
        birdsCount+=birdsPerDay[i]
    }
	return birdsCount
}

 BirdsInWeek returns the total bird count by summing
 only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	count = 0
	weekCount = 0
	for i = 1; i  week; i++ {
		weekCount += 7
	}
	for i = weekCount; i  weekCount+7; i++ {
		count += birdsPerDay[i]
	}
	return count
}

 FixBirdCountLog returns the bird counts after correcting
 the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i = 0; i  len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
    
}
