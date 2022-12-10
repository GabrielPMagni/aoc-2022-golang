package functions_test

import (
	"aoc/go2022/functions"
	"testing"
)

func compareAllLines(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i := 0; i <= len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func TestGetLinesFunction(t *testing.T) {
	filePath := "test_input.text"

	testSuite := []struct {
		name          string
		expectedLines []string
	}{
		{
			name: "Test base lines extraction",
			expectedLines: []string{
				"4514",
				"8009",
				"6703",
				"1811",
				"4881",
				"3905",
				"3933",
				"9436",
				"4332",
				"",
				"3059",
				"15715",
				"11597",
				"10625",
				"8486",
			},
		},
	}

	for _, ts := range testSuite {
		gotLines := functions.GetFileLines(filePath)

		if compareAllLines(gotLines, ts.expectedLines) {
			t.Errorf("got unexpected lines -> got: [%+v] , expected: [%+v]", gotLines, ts.expectedLines)
		}
	}
}

// go test ./...
// go test ./functions/...
