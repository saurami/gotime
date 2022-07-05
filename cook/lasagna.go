package cook

const ovenTime = 40

func remainingOvenTime(actual int8) int8 {
	var remaining int8 = ovenTime - actual
	return remaining
}

func preparationTime(numLayers int8) int8 {
	return numLayers * 2
}

func elapsedTime(numLayers int8, ovenMins int8) int8 {
	var cooked_for = preparationTime(numLayers) + ovenMins
	return cooked_for
}
