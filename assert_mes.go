package mes

import (
	"reflect"
	"testing"
)

func MatchExactly(t *testing.T, expected, actual interface{}) {
	if !exactMatch(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func exactMatch(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return false
	}

	if reflect.TypeOf(expected).Kind() != reflect.Slice || reflect.TypeOf(actual).Kind() != reflect.Slice {
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
