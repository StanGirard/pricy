package helpers

import "testing"

func TestFloat2Decimal(t *testing.T) {
	if float2Decimal(1.2345) != "1.23" {
		t.Errorf("Expected 1.23, got %s", float2Decimal(1.2345))
	}
}
