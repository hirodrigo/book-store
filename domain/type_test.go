package domain

import (
	"testing"
)

func TestFictionBookType(t *testing.T) {
	expected := Fiction
	actual := FromString("Non-Fiction")

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
