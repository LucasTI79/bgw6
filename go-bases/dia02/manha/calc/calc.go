package calc

import "errors"

const (
	Sum  = "+"
	Sub  = "-"
	Mult = "*"
	Div  = "/"
)

var ErrCannotDivideByZero = errors.New("can not divide by zero")

func CalculateCurrying(operator string) func(value1, value2 float64) float64 {
	switch operator {
	case Sum:
		return opSummatory
	case Sub:
		return opSubtraction
	case Mult:
		return opMultiplication
	case Div:
		return opDivision
	}

	return nil
}

func Calculate(operator string, values ...float64) float64 {
	switch operator {
	case Sum:
		return operationsOrchestrator(values, opSummatory)
	case Sub:
		return operationsOrchestrator(values, opSubtraction)
	case Mult:
		return operationsOrchestrator(values, opMultiplication)
	case Div:
		return operationsOrchestrator(values, opDivision)
	}

	return 0
}

func operationsOrchestrator(values []float64, operation func(value1, value2 float64) float64) float64 {
	var result float64
	for _, value := range values {
		result = operation(result, value)
	}

	return result
}

func opSummatory(value1, value2 float64) float64 {
	return value1 + value2
}

func opSubtraction(value1, value2 float64) float64 {
	return value1 - value2
}

func opMultiplication(value1, value2 float64) float64 {
	return value1 * value2
}

func opDivision(value1, value2 float64) float64 {
	if value2 == 0 {
		return 0
	}

	return value1 / value2
}

func Division(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, ErrCannotDivideByZero
	}
	return numerator / denominator, nil
}
