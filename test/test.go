package main

import (
		"gBay"
		"fmt"
	)

func main(){
	fmt.Println("Test")
	snd := gBay.SliceNDice{FileName: "data.csv" , SplitRatio: 1}
	snd.ReadCsv()
	a, _ := snd.Split()
	fmt.Println(snd.ByClass(a))
	fmt.Println(gBay.GaussPDF(71.5, 73, 6.2))
}
