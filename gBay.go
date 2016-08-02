package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func handleErr(err error){
	if err != nil{
		panic(err)
	}
}

type sliceNDice struct{
	fileName string
	splitRatio float64
}

func(s sliceNDice) readCsv() [][]string {
	var dataset [][]string
	file, err := os.Open(s.fileName)
	handleErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	
	rawData, err := reader.ReadAll()
	handleErr(err)

	for _, vals := range rawData{
		dataset = append(dataset,vals)
	}
	return dataset
}

func main(){
	snd := sliceNDice{fileName: "data.csv" , splitRatio: 0.75}
	dataset := snd.readCsv()
	fmt.Println(dataset)
}
