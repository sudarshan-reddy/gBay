package gBay

func CalcProbs(summaries map[int][][]float64 ,
		vector []float64) map[int]float64{
	prob := make(map[int]float64)
	for classVal, classSum := range summaries{
		prob[classVal] = 1
		for  k , v := range classSum{
			mean , stdev := v[0] , v[1]
			x := vector[k]
			prob[classVal] *= GaussPDF(x, mean, stdev)		
		}

	}
	return prob
}

func Predict(summaries map[int][][]float64 ,
		vector []float64) int{
	bestProb, bestLabel := -1.0 , -1
	probabilities := CalcProbs(summaries, vector)
	for classVal, probability := range probabilities{
		if bestLabel == -1 || probability > bestProb{
			bestProb = probability
			bestLabel = classVal
		}
	}
	return bestLabel
}

func getPredictions(summaries map[int][][]float64 ,
		vector []float64){
		
}
