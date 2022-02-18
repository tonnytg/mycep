package location

import (
	"testing"
)

func TestSetCEP(t *testing.T) {
	location := NewLocation()
	location.SetCEP("12345-678")
	if location.CEP != "12345-678" {
		t.Errorf("CEP was not set correctly")
	}
}
