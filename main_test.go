package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		tier      string
		expected  int
	}
	tests := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"invalid", 0},
			{"", 0},
		}...)
	}

	for _, test := range tests {
		if output := getMontlyPrice(
			test.tier,
		); output != test.expected {
			t.Errorf(
				"Test Failed: (%v) -> expected: %v actual: %v",
				test.tier, test.expected, output)
		} else {
			fmt.Printf("Test Passed: (%v) -> expected: %v actual: %v\n",
		test.tier, test.expected, output)
		}
	}
}

var withSubmit = true