package match

import (
	"reflect"
	"testing"
)

// Exactly compares two slices and returns true if they are _exactly_ the same.
// This means that the slices must have the same length and the same elements in the same order.
func Exactly(t *testing.T, expected, actual interface{}) {
	if !matchExactly(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

// Elements compares two slices and returns true if they contain the same elements.
// This means that the slices must have the same length and the same elements, but not necessarily in the same order.
func Elements(t *testing.T, expected, actual interface{}) {
	if !matchElements(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func matchExactly(expected, actual interface{}) bool {
	if !validType(expected, actual) {
		return false
	}

	s1 := reflect.ValueOf(expected)
	s2 := reflect.ValueOf(actual)

	if s1.Len() != s2.Len() {
		return false
	}

	for i := 0; i < s1.Len(); i++ {
		if s1.Index(i).Interface() != s2.Index(i).Interface() {
			return false
		}
	}

	return true
}

func matchElements(expected, actual interface{}) bool {
	if !validType(expected, actual) {
		return false
	}

	s1 := reflect.ValueOf(expected)
	s2 := reflect.ValueOf(actual)

	if s1.Len() != s2.Len() {
		return false
	}

	counter := make(map[interface{}]int)
	for i := 0; i < s1.Len(); i++ {
		element := s1.Index(i).Interface()
		counter[element]++
	}

	for i := 0; i < s2.Len(); i++ {
		element := s2.Index(i).Interface()
		if counter[element] == 0 {
			return false
		}
		counter[element]--
	}

	return true
}

func validType(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return false
	}

	if reflect.TypeOf(expected).Kind() != reflect.Slice || reflect.TypeOf(actual).Kind() != reflect.Slice {
		return false
	}

	return true
}
