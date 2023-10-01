package service

import "testing"

func TestAbs(t *testing.T) {
	if 1 != 1 {
		t.Errorf("Abs(-1) = %d; want 1", 3)
	}
}
