package wtconv

// KgToLb Умножаем значение массы на 2.205
func KgToLb(kg Kg) Lb {
	return Lb(kg * 2.205)
}

// LbToKg Разделим значение массы на 2.205
func LbToKg(lb Lb) Kg {
	return Kg(lb / 2.205)
}
