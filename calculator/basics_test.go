package calculator

import "testing"

func TestAdd(t *testing.T) {

	t.Run("adding two positive numbers", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4
		
		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})
	
	t.Run("adding two negative numbers", func(t *testing.T) {
		sum := add(-3, -4)
		expected := -7

		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})

	t.Run("adding one positive and one negative integer", func(t *testing.T) {
		sum := add(1, -3)
		expected := -2

		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})
}


func TestSubtract(t *testing.T) {

	t.Run("subtracting two positive numbers", func(t *testing.T) {
		sum := subtract(2, 2)
		expected := 0
		
		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})
	
	t.Run("subtracting two negative numbers", func(t *testing.T) {
		sum := subtract(-3, -4)
		expected := 1

		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})

	t.Run("subtracting a positive integer from a negative integer", func(t *testing.T) {
		sum := subtract(1, -3)
		expected := 4

		if sum != expected {
			t.Errorf("Expected %d; but got %d", expected, sum)
		}
	})
}
