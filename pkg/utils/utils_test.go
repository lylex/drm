package utils

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	result, err := Marshal(group)
	expect := `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`

	if result != expect {
		t.Errorf("failed to marshal object, expect %v, got %v", expect, result)
	}
	if err != nil {
		t.Errorf("unexpected error occurred: %v", err)
	}
}

func TestUnmarshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{}
	str := `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`

	err := Unmarshal(str, &group)
	expect := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	if group.Name != expect.Name || group.ID != expect.ID {
		t.Errorf("failed to unmarshal string to object: expect %v, got %v", expect, group)
	}
	if err != nil {
		t.Errorf("unexpected error occurred: %v", err)
	}
}
func TestGenerateRandString(t *testing.T) {
	var testcases = []struct {
		n         int
		resultLen int
	}{
		{1, 1},
		{6, 6},
	}

	for _, testcase := range testcases {
		result := GenerateRandString(testcase.n)
		if len(result) != testcase.resultLen {
			t.Errorf("Generate radom string failed, expect %d, got %d", testcase.resultLen, len(result))
		}
	}

	// We decide to ignore the tiny probability.
	if GenerateRandString(6) == GenerateRandString(6) {
		t.Errorf("Generate radom string should not be the same")
	}
}
