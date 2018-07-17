package measureconv

import (
	"testing"
)

func TestMeasureConv(t *testing.T) {
	var i Inches = 12.0
	var c Centimeters = 6.0

	itc := IToCm(i)
	cti := CmToI(c)

	if itc != 30.48 {
		t.Errorf("Expected conversion of '30.48' but got %v", itc)
	}

	if cti != 2.34 {
		t.Errorf("Expected conversion of '2.34' but got %v", cti)
	}

}
