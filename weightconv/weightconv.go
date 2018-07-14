package weightconv

type Pounds float64
type Kilograms float64

func LbsToKgs(p Pounds) Kilograms {
	return Kilograms(p * 0.4536)
}

func KgsToLbs(k Kilograms) Pounds {
	return Pounds(k / 0.4536)
}
