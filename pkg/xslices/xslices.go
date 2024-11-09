package xslices

// Map applies a given function to each element of the input slice and returns a new slice with the results.
// Useful for transforming elements in a slice.
func Map[T any, U any](input []T, transform func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}

// Filter returns a new slice containing only elements from the input slice that satisfy the predicate function.
// Useful for selecting elements based on a condition.
func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces a slice to a single value by applying the combine function to each element and an accumulator.
// Useful for aggregating or summarizing slice elements.
func Reduce[T any, U any](input []T, initial U, combine func(U, T) U) U {
	accumulator := initial
	for _, v := range input {
		accumulator = combine(accumulator, v)
	}
	return accumulator
}

// Unique returns a new slice with only the unique elements from the input slice, preserving the original order.
// Useful for removing duplicates from a slice.
func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	var result []T
	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// UniqueFunc returns a new slice containing only unique elements from the input slice,
// using a keyExtractor function to determine uniqueness.
// The keyExtractor function generates a key for each element, and elements with the same key are considered duplicates.
// This is useful for filtering out duplicate elements based on specific properties.
func UniqueFunc[T any, K comparable](input []T, keyExtractor func(T) K) []T {
	seen := make(map[K]struct{})
	var result []T
	for _, v := range input {
		key := keyExtractor(v)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Any returns true if any element in the slice satisfies the predicate function; otherwise, it returns false.
// Useful for checking if at least one element matches a condition.
func Any[T any](input []T, predicate func(T) bool) bool {
	for _, v := range input {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate function; otherwise, it returns false.
// Useful for verifying that all elements match a condition.
func All[T any](input []T, predicate func(T) bool) bool {
	for _, v := range input {
		if !predicate(v) {
			return false
		}
	}

	return true
}
