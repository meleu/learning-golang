package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15
		assertEqualSum(t, actual, expected, numbers)
	})
}

func assertEqualSum(t testing.TB, actual, expected int, givenNumbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf(
			"\n  actual: %d\nexpected: %d\ngiven: %v",
			actual, expected, givenNumbers,
		)
	}
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 1, 9})
	expected := []int{3, 10}

	if !slices.Equal(actual, expected) {
		t.Errorf("\n  actual: %v\nexpected: %v", actual, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	assertEqualSums := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !slices.Equal(actual, expected) {
			t.Errorf("\n  actual: %v\nexpected: %v", actual, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 1, 9})
		expected := []int{2, 10}
		assertEqualSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		assertEqualSums(t, actual, expected)
	})
}
