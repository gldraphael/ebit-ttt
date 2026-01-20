package core

import (
	"testing"
)

var validateCoordsTests = []struct {
	x            int
	y            int
	expectsError bool
}{
	{ 0,  0, false},
	{ 2,  2, false},
	{ 3,  3, true },
	{-1, -1, true },
}

func TestValidateCoords(t *testing.T) {
	for _, testCase := range validateCoordsTests {
		got := validateCoords(testCase.x, testCase.y)
		if (got != nil) != testCase.expectsError {
			t.Errorf(
				"Test case failed for validateCoords(%d, %d). \nExpected: %t \nGot: Error: %s",
				testCase.x, testCase.y, testCase.expectsError, got.Error())
		}
	}
}
