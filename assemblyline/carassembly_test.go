package assemblyline

import "testing"

func TestWorkingCars(t *testing.T) {
	t.Run("working cars per hour", func(t *testing.T) {
		got := calculateWorkingCarsPerHour(1547, 90)
		want := 1392.3

		if got != want {
			t.Errorf("incorrect number of cars produced per hour; got %v, want %v", got, want)
		}
	})

	t.Run("working cars per minute", func(t *testing.T) {
		got := calculateWorkingCarsPerMinute(1105, 90)
		want := 16

		if got != want {
			t.Errorf("incorrect number of cars produced per minute; got %v, want %v", got, want)
		}
	})
}

func TestProductionCost(t *testing.T) {
	var tests = []struct {
		description string
		got uint32
		want uint32
	}{
		{"individual cars only", calculateProductionCost(6),  60000},
		{"two groups of cars",   calculateProductionCost(21), 200000},
		{"three groups of cars", calculateProductionCost(37), 355000},
	}

	for _, test := range tests {
		t.Logf("testing for %s", test.description)
		if test.got != test.want {
			t.Errorf("Issue calculating production cost; got %v, watn %v", test.got, test.want)
		}
	}
}
