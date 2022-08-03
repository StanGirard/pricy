package helpers

import (
	"reflect"
	"testing"
)

func TestConvertStringToInterface(t *testing.T) {
	array := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	expected := [][]interface{}{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	actual := ConvertStringToInterface(array)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
