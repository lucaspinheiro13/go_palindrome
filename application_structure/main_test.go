package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		input          string
		expectedResult bool
	}

	testCases := []Case{
		{
			input:          "Ana",
			expectedResult: true,
		},
		{
			input:          "mamam",
			expectedResult: true,
		},
		{
			input:          "Lucas",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.expectedResult {
			t.Error("Fail", testCase)
		}
	}
}
