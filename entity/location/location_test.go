package location

import (
	"testing"
)

func TestSetCEP(t *testing.T) {
	location := NewLocation()
	location.SetCEP("12345678")
	if location.CEP != "12345678" {
		t.Errorf("CEP was not set correctly")
	}

	if location.Street != "" {
		t.Errorf("Street was not set correctly")
	}
}
