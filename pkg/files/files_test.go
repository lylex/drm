package files

import (
	"testing"
)

func TestIsAbsolutePath(t *testing.T) {
	var testcases = []struct {
		path           string
		expectedResult bool
	}{
		{"/dir", true},
		{"dir", false},
		{"/dir1/dir2", true},
		{"dir1/dir2", false},
		{"", false},
	}

	for _, testcase := range testcases {
		r := IsAbsolutePath(testcase.path)
		if r != testcase.expectedResult {
			t.Errorf("Judge \"%s\" failed, expect %v, got %v", testcase.path, testcase.expectedResult, r)
		}
	}
}
