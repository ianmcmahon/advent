package intcode

import "testing"

func TestParamModes(t *testing.T) {
	if a, b := iioParamModes(110); a != IMMEDIATE || b != IMMEDIATE {
		t.Errorf("%d => %d, %d", 110, a, b)
	}
	if a, b := iioParamModes(0); a != POSITION || b != POSITION {
		t.Errorf("%d => %d, %d", 0, a, b)
	}
	if a, b := iioParamModes(10); a != POSITION || b != IMMEDIATE {
		t.Errorf("%d => %d, %d", 10, a, b)
	}
	if a, b := iioParamModes(100); a != IMMEDIATE || b != POSITION {
		t.Errorf("%d => %d, %d", 100, a, b)
	}

	// can't check the output position without catching panic
}
