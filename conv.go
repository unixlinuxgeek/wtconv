package wtconv

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / 2.2046226218)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg * 2.2046226218)
}
