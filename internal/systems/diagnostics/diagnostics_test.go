package diagnostics

import "testing"

func TestSimpleDiagnostics_CalculatePowerUsage(t *testing.T) {
	type testDef struct {
		name   string
		bits   int
		input  []string
		output PowerUsage
	}

	tests := []testDef{
		{
			name: "Should generate expected values according to the example input",
			bits: 5,
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			output: PowerUsage{
				Gamma:   22,
				Epsilon: 9,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := NewDiagnostics(tt.bits)

			summary, err := diag.CalculatePowerUsage(tt.input...)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.output.String() != summary.String() {
				t.Errorf("expected different output summary. Expected: %v, Got: %v", tt.output, summary)
				return
			}
		})
	}
}

func TestSimpleDiagnostics_CalculateLifeSupportRating(t *testing.T) {
	type testDef struct {
		name   string
		bits   int
		input  []string
		output LifeSupportRating
	}

	tests := []testDef{
		{
			name: "Should generate expected values according to the example input",
			bits: 5,
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",

				// 10110
				// 10111
			},
			output: LifeSupportRating{
				OxygenGenerator: 23,
				CO2Scrubber:     10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag := NewDiagnostics(tt.bits)

			summary, err := diag.CalculateLifeSupportRating(tt.input...)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.output.String() != summary.String() {
				t.Errorf("expected different output summary. Expected: %v, Got: %v", tt.output, summary)
				return
			}
		})
	}
}
