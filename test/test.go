package main

import (
		"gBay"
		"fmt"
	)

func main(){
	fmt.Println("Test")
	snd := gBay.SliceNDice{FileName: "data.csv" , SplitRatio: 0.75}
	snd.ReadCsv()
	training , test := snd.Split()
	//prepare model
	summaries := snd.ByClass(training)
	//predictions	
	sums := map[int][][]float64{
		0 : [][]float64{[]float64{1 , 0.5}}, 
		1 : [][]float64{[]float64{20 , 5.0}}, 
	}

	vector := [][]float64{
		[]float64{1.1 , 0}, 
		[]float64{19.1 , 0},
	}

	fmt.Println(gBay.GetPredictions(sums, vector))
	fmt.Println(gBay.GetPredictions(summaries, test))
}
