package vehiclepurchase

import "testing"

func TestVehicleLicense(t *testing.T) {
	var tests = []struct {
		description string
		got bool
		want bool
	}{
		{"car", needsLicense("car"), true},
		{"truck", needsLicense("truck"), true},
		{"bike", needsLicense("bike"), false},
	}

	for _, testcase := range tests {
		t.Logf("testing license requirement for %s", testcase.description)
		if testcase.got != testcase.want {
			t.Errorf("Incorrect license requirement ... got %v, want %v", testcase.got, testcase.want)
		}
	}
}

func TestVehicleSelection(t *testing.T) {
	var tests = []struct {
		description string
		got string
		want string
	}{
		{
			description: "Wuling Hongguang and Toyota Corolla",
			got: chooseVehicle("Wuling Hongguang", "Toyota Corolla"),
			want: "Toyota Corolla",
		},
		{
			description: "Volkswagen Beetle and Volkswagen Golf",
			got: chooseVehicle("Volkswagen Beetle", "Volkswagen Golf"),
			want: "Volkswagen Beetle",
		},
	}

	for _, testcase := range tests {
		t.Logf("testing vehicle selection between %s", testcase.description)
		if testcase.got != testcase.want {
			t.Errorf("Error making selection ... got %v, want %v", testcase.got, testcase.want)
		}
	}
}
