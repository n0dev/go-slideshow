package utils

import "testing"

func TestMod(t *testing.T) {
	if Mod(0, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(1, 15) != 1 {
		t.Errorf("Error")
	}
	if Mod(15, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(-1, 15) != 14 {
		t.Errorf("Error")
	}
	if Mod(-15, 15) != 0 {
		t.Errorf("Error")
	}
	if Mod(-16, 15) != 14 {
		t.Errorf("Error")
	}
}
