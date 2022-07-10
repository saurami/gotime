package assemblyline

func calculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCars := (float64(productionRate) * successRate) / 100
	return workingCars
}

func calculateWorkingCarsPerMinute(carsProducedPerHour int, successRate float64) int {
	workingCars := calculateWorkingCarsPerHour(carsProducedPerHour, successRate)
	workingCars = workingCars / 60
	return int(workingCars)
}

func calculateProductionCost(numCars uint32) uint32 {
	var carCost uint32 = 10000
	var carGroupCost uint32 = 95000  // group of 10 cars

	if numCars >= 10 {
		individualCars := numCars % 10
		groupsOfTen := numCars / 10
		return (groupsOfTen * carGroupCost) + (individualCars * carCost)
	} else {
		return numCars * carCost
	}
}
