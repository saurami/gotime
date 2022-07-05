package gamequest

import "testing"

func TestAttackIfKnightAwake(t *testing.T) {
	tests := []struct {
		description string
		input bool
		expected bool
	}{
		{
			description: "fast attack cannot be made",
			input: canFastAttack(true),
			expected: false,
		},
		{
			description: "fast attack can be made",
			input: canFastAttack(false),
			expected: true,
		},
	}

	for _, testcase := range tests {
		t.Logf("testing whether %s", testcase.description)
		if testcase.input != testcase.expected {
			t.Errorf("trouble making fast attack ... got %v, wanted %v", testcase.input, testcase.expected)
		}
	}
}

func TestSpyOnGroup(t *testing.T) {
	tests := []struct {
		description string
		input bool
		expected bool
	}{
		{
			description: "archer awake; knight and prisoner asleep",
			input: canSpy(false, true, false),
			expected: true,
		},
		{
			description: "everybody's sleeping; spying a waste of time ...",
			input: canSpy(false, false, false),
			expected: false,
		},
	}

	for _, testcase := range tests {
		t.Logf("testing spy if %s", testcase.description)
		if testcase.input != testcase.expected {
			t.Errorf("issue spying on group ... got %v, wanted %v", testcase.input, testcase.expected)
		}
	}
}

func TestSignalToPrisoner(t *testing.T) {
	tests := []struct {
		description string
		input bool
		expected bool
	}{
		{
			description: "archer is asleep and prisoner is awake",
			input: canSignalPrisoner(false, true),
			expected: true,
		},
		{
			description: "archer and prisoner both awake",
			input: canSignalPrisoner(true, true),
			expected: false,
		},
		{
			description: "archer and prisoner both asleep",
			input: canSignalPrisoner(false, false),
			expected: false,
		},
		{
			description: "archer is awake and prisoner asleep",
			input: canSignalPrisoner(true, true),
			expected: false,
		},
	}

	for _, testcase := range tests {
		t.Logf("Testing if signal can be made if %s", testcase.description)
		if testcase.input != testcase.expected {
			t.Errorf("trouble making signal ... got %v, wanted %v", testcase.input, testcase.expected)
		}
	}
}

func TestCanFreePrisoner(t *testing.T) {
	tests := []struct {
		description string
		input bool
		expected bool
	}{
		{
			description: "archer is awake, knight and prisoner are asleep, and dog is absent",
			input: canFreePrisoner(false, true, false, false),
			expected: false,
		},
	}

	for _, testcase := range tests {
		t.Logf("testing if prisoner can be freed when %s", testcase.description)
		if testcase.input != testcase.expected {
			t.Errorf("unable to free prisoner ... got %v, wanted %v", testcase.input, testcase.expected)	
		}
	}
}