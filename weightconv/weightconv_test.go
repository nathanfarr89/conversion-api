package weightconv

import (
	"testing"
)

func TestWeightConv(t *testing.T) {
	var p Pounds = 245
	var k Kilograms = 100

	ptk := LbsToKgs(p)
	ktp := KgsToLbs(k)

	if ptk != 111.132 {
		t.Errorf("Expected conversion of '111.132' but got %v", ptk)
	}

	if ktp != 220.4585537918871 {
		t.Errorf("Expected conversion of '220.4585537918871' but got %v", ktp)
	}

}
