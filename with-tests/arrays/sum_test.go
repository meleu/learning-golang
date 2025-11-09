package main

import "testing"

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
