package gBay

import (
		"math"
		//"fmt"
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

func StDev(dataset []float64) float64{
	var variance float64
	for _ , values := range dataset{
		variance+=(math.Pow((Mean(dataset)-values),2))/float64(len(dataset)-1)
	}
	return math.Sqrt(variance)
}

func Summarize(d [][]float64) [][]float64 {
	for i := 0; i < len(d) - 1 ; i++{
		if len(d[i]) != len(d[i + 1]){
			return nil //Too lazy to return errors now. Marked as todo. 
		}
	}
	
	retD := make([][]float64, len(d[0]))
	zipA := make([][]float64, len(d[0]))
	
	for i := 0; i < len(d[0]) ; i++{
		for _, vals := range d{
			zipA[i] = append(zipA[i] , vals[i])
		}
		retD[i] = []float64{Mean(zipA[i]),StDev(zipA[i])}
	}

	return retD
}
