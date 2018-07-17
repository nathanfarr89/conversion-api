package tempconv

import (
	"testing"
)

func TestTempConv(t *testing.T) {
	var c Celsius = 30.0
	var f Fahrenheit = 80.0
	var k Kelvin = 40.0

	ftc := FToC(f)
	ktc := KtC(k)
	ctf := CToF(c)
	ktf := KtF(k)
	ftk := FtK(f)
	ctk := CtK(c)

	if ftc != 26.666666666666668 {
		t.Errorf("Expected conversion of '' but got %v", ftc)
	}

	if ktc != -233.14999999999998 {
		t.Errorf("Expected conversion of '-233.14999999999998' but got %v", ktc)
	}

	if ctf != 86 {
		t.Errorf("Expected conversion of '86' but got %v", ctf)
	}

	if ktf != -387.67 {
		t.Errorf("Expected conversion of '-387.67' but got %v", ktf)
	}

	if ftk != 299.8166666666667 {
		t.Errorf("Expected conversion of '299.8166666666667' but got %v", ftk)
	}

	if ctk != 303.15 {
		t.Errorf("Expected conversion of '303.15' but got %v", ctk)
	}

}
