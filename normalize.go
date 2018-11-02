package normalize

//Normalize returns a list of normalized values
func Normalize(list []float64) (result []float64) {
	max := Max(list)
	min := Min(list)
	result = make([]float64, len(list))
	for index, value := range list {
		result[index] = normalize(value, min, max)
	}
	return result
}

func normalize(value float64, min float64, max float64) float64 {
	if min != max {
		return (value - min) / (max - min)
	}
	return 1.0
}

//Min calculates the maximum value in a slice
func Min(list []float64) (min float64) {
	min = list[0]
	for _, number := range list {
		if number < min {
			min = number
		}
	}
	return min
}

//Max calculates the maximum value in a slice
func Max(list []float64) (max float64) {
	max = list[0]
	for _, number := range list {
		if number > max {
			max = number
		}
	}
	return max
}
