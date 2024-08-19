package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T)  {
	
	t.Run("Collection of fixed sixe",func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers)
		want:= 15
	
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T)  {
	got := SumAll([] int {1,2,3}, [] int {1,2,3}, [] int {1,2,3})
	expected := [] int {6,6,6}

	// if !reflect.DeepEqual(got,expected) {
	// 	t.Errorf("got %v want %v", got, expected)
	// }

	if !slices.Equal(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}