package main

import (
		"gBay"
		"fmt"
	)

func main(){
	fmt.Println("Test")
	snd := gBay.SliceNDice{FileName: "data.csv" , SplitRatio: 0.75}
	snd.ReadCsv()
	training , _ := snd.Split()
	//prepare model
	summaries := snd.ByClass(training)
	fmt.Println(summaries)
	vector := []float64{1.1, 0}
	//predictions	
	sums := map[int][][]float64{
		0 : [][]float64{[]float64{1 , 0.5}}, 
		1 : [][]float64{[]float64{20 , 5.0}}, 
	}
	fmt.Println(sums)
	fmt.Println(gBay.CalcProbs(sums, vector))
	fmt.Println(gBay.CalcProbs(summaries, vector))
}
