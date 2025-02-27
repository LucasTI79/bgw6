package calc

func Operations(
	value1,
	value2 float64,
) (
	summatory float64,
	subtraction float64,
	multiplication float64,
	division float64,
) {
	summatory = value1 + value2
	subtraction = value1 - value2
	multiplication = value1 * value2

	if value2 != 0 {
		division = value1 / value2
	}

	return summatory, subtraction, multiplication, division
}
