package assert

import "fmt"

// Equal checks if two values are equal and panics with a message if they are not.
// Useful for testing value equality.
func Equal[T comparable](expected, actual T, message string) {
	if expected != actual {
		panic(fmt.Sprintf("Assertion failed: %s - expected: %v, got: %v", message, expected, actual))
	}
}

// NotNil checks if a given value is not nil and panics with a message if it is nil.
// Useful for testing that a pointer or interface contains a value.
func NotNil(value any, message string) {
	if value == nil {
		panic(fmt.Sprintf("Assertion failed: %s - value should not be nil", message))
	}
}

// True checks if a given condition is true and panics with a message if it is false.
// Useful for asserting conditions in tests.
func True(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Assertion failed: %s - expected true, got false", message))
	}
}

// False checks if a given condition is false and panics with a message if it is true.
// Useful for asserting negative conditions in tests.
func False(condition bool, message string) {
	if condition {
		panic(fmt.Sprintf("Assertion failed: %s - expected false, got true", message))
	}
}

// Nil checks if a given value is nil and panics with a message if it is not nil.
// Useful for testing that a pointer or interface is nil.
func Nil(value any, message string) {
	if value != nil {
		panic(fmt.Sprintf("Assertion failed: %s - expected nil, got %v", message, value))
	}
}

// Func takes a function that returns a boolean and checks if it returns true.
// Panics with a message if the function returns false.
// Useful for asserting custom conditions encapsulated within a function.
func Func(conditionFunc func() bool, message string) {
	if !conditionFunc() {
		panic(fmt.Sprintf("Assertion failed: %s - function returned false", message))
	}
}
