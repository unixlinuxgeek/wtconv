package wtconv

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / lb)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg * lb)
}
