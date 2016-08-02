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

func(s sliceNDice) readCsv() []float64 {
	file, err := os.Open(s.fileName)
	handleErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	
	rawData, err := reader.ReadAll()
	handleErr(err)

	for _, vals := range rawData{
		fmt.Println(vals)
	}
	return []float64{0.1}
}

func main(){
	snd := sliceNDice{fileName: "data.csv" , splitRatio: 0.75}
	fmt.Println(snd.readCsv())
}
