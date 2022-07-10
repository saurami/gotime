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
