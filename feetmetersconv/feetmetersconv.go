package feetmetersconv

type Feet float64
type Meters float64

func FtM(f Feet) Meters {
	return Meters(f * 0.3048)
}

func MtF(m Meters) Feet {
	return Feet(m / 0.3048)
}
