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

