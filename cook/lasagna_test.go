package cook

import "testing"

func TestLasagnaCooking(t *testing.T) {
	tests := []struct {
		description string
		input       int8
		expected    int8
	}{
		{
			description: "remaining oven time in minutes",
			input:       remainingOvenTime(30),
			expected:    10,
		},
		{
			description: "preparation time in minutes",
			input:       preparationTime(2),
			expected:    4,
		},
		{
			description: "elapsed cooking time in minutes",
			input:       elapsedTime(3, 20),
			expected:    26,
		},
	}

	for _, testcase := range tests {
		t.Logf("Testing for %s", testcase.description)
		if testcase.input != testcase.expected {
			t.Errorf("Lasagna doesn't taste good ... got %v; expected %v", testcase.input, testcase.expected)
		}
	}
}
