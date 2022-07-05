package gamequest

func canFastAttack(knightAwake bool) bool {
	if knightAwake == true {
		return false
	} else {
		return true
	}
}

func canSpy(knightAwake, archerAwake, prisonerAwake bool) bool {
	if knightAwake == true || archerAwake == true || prisonerAwake == true {
		return true
	}
	return false
}

func canSignalPrisoner(archerAwake bool, prisonerAwake bool) bool {
	if prisonerAwake == true && archerAwake == false {
		return true
	} else {
		return false
	}
}

func canFreePrisoner(knightAwake, archerAwake, prisonerAwake bool, dogPresent bool) bool {
	if dogPresent == true {
		if archerAwake == false {
			return true
		} else {
			return false
		}
	} else if dogPresent == false {
		if knightAwake == false && archerAwake == false && prisonerAwake == true {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
