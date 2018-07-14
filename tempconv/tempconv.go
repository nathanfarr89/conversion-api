package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

func KtC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func KtF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
