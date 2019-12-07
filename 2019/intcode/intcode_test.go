package intcode

import "testing"

func TestParamModes(t *testing.T) {
	if a, b, _ := iioParamModes(11); a != IMMEDIATE || b != IMMEDIATE {
		t.Errorf("%d => %d, %d", 11, a, b)
	}
	if a, b, _ := iioParamModes(0); a != POSITION || b != POSITION {
		t.Errorf("%d => %d, %d", 0, a, b)
	}
	if a, b, _ := iioParamModes(10); a != POSITION || b != IMMEDIATE {
		t.Errorf("%d => %d, %d", 10, a, b)
	}
	if a, b, _ := iioParamModes(1); a != IMMEDIATE || b != POSITION {
		t.Errorf("%d => %d, %d", 1, a, b)
	}

	// can't check the output position without catching panic
}
