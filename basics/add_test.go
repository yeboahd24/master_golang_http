package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}
}

// TableDrivenTest is a test that runs a function with different inputs and
func TestTableDrivenTest(t *testing.T) {
	testData := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}
	for _, td := range testData {
		if got := Add(td.a, td.b); got != td.want {
			t.Errorf("Add(%d, %d) = %d, want %d", td.a, td.b, got, td.want)
		}
	}
}
