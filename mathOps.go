package gBay

import (
		"math"
		"fmt"
	)

//parametric polymorphism here later
func Sum(dataset []float64) float64{
	var sum float64
	for _, vals := range dataset{
		sum += vals
	}
	return sum
}

func Mean(dataset []float64) float64{
	return Sum(dataset) / float64(len(dataset))
}

func stdDev(dataset []float64) float64{
	var variance float64
	for _ , values := range dataset{
		variance+=(math.Pow((Mean(dataset) - values),2))/float64(len(dataset))
	}
	return math.Sqrt(variance)

}

