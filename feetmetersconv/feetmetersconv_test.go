package feetmetersconv

import (
	"testing"
)

func TestFeetmetersConv(t *testing.T) {
	var f Feet = 6
	var m Meters = 3

	ftm := FtM(f)
	mtf := MtF(m)

	if ftm != 1.8288000000000002 {
		t.Errorf("Expected conversion of '1.8288000000000002' but got %v", ftm)
	}

	if mtf != 9.84251968503937 {
		t.Errorf("Expected conversion of '9.84251968503937' but got %v", mtf)
	}

}
