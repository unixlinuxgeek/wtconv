package wtconv

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / Lb)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg * Lb)
}
