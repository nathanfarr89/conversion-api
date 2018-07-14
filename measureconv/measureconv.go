package measureconv

type Inches float64
type Centimeters float64

func IToCm(i Inches) Centimeters {
	return Centimeters(i * 2.54)
}

func CmToI(c Centimeters) Inches {
	return Inches(c * .39)
}
