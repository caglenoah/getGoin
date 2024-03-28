package main
import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costPerSend,
		numLastMonth,
		numThisMonth,
		expected int
	}
	tests := []testCase{
		{2, 89, 102, 26},
		{2, 98, 104, 12},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{3, 50, 40, -30},
			{3, 60, 60, 0},
		}...)
	}
	for _, test := range tests {
		if output := monthlyBillIncrease(
			test.costPerSend,
			test.numLastMonth,
			test.numThisMonth,
		); output != test.expected {
			t.Error(
				"Test Failed: (%v, %v, %v) -> expected: %v actual: %v\n",
				test.costPerSend,
				test.numLastMonth,
				test.numThisMonth,
				test.expected,
				output,
			)
		} else {
			fmt.Printf("Test Pased: (%v, %v, %v) -> expected: %v actual: %v\n",
		test.costPerSend,
		test.numLastMonth,
		test.numThisMonth,
		test.expected,
		output,
		)
		}
	}
}


withSubmit = true
