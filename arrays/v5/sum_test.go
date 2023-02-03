package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

//no need for two tests here because if the program works for slices, it would work for any size of slices.
func TestSumALl(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{3, 6})
	want := []int{4, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestTailSumAll(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}

	t.Run("adding non zero tail items", func(t *testing.T) {

		got := TailSumAll([]int{1, 3}, []int{2, 4})
		want := []int{3, 4}

		checkSums(t, got, want)
	})
	t.Run("adding possible zero tail items", func(t *testing.T) {

		got := TailSumAll([]int{}, []int{2, 4})
		want := []int{0, 4}
		checkSums(t, got, want)

	})
}
