package gBay

func CalcProbs(summaries map[int][][]float64 ,
		vector []float64) map[int]float64{
	prob := make(map[int]float64)
	for classVal, classSum := range summaries{
		prob[classVal] = 1
		for  _ , v := range classSum{
			mean , stdev := v[0] , v[1]
			x := vector[0]
			prob[classVal] *= GaussPDF(x, mean, stdev)		
		}

	}
	return prob
}
	
