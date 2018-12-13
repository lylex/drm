package utils

import "testing"

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
		t.Errorf("unexpected error occured: %v", err)
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
		t.Errorf("unexpected error occured: %v", err)
	}
}
