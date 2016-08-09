package gBay

func CalcProbs(summaries map[int][][]float64 , vector []float64) map[int]float64{
	prob := make(map[int]float64)
	for classVal, classSum := range summaries{
		prob[classVal] = 1
		for  k , v := range classSum{
			if k == len(classSum) - 1{
				break
			}
			mean , stdev := v[0] , v[1]
			x := vector[0]
			prob[classVal] *= GaussPDF(x, mean, stdev)		
		}

	}
	return prob
}
	
