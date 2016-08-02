package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

//put these in a separate file later
//or not
func handleErr(err error){
	if err != nil{
		panic(err)
	}
}

func str2flt(ip []string) []float64{
	var retfloat []float64
	for _ , val := range ip{
		if n, err := strconv.ParseFloat(val,64);err == nil{
			retfloat = append(retfloat, n)
		}
	}
	return retfloat
}

type sliceNDice struct{
	fileName string
	splitRatio float64
}

func(s sliceNDice) readCsv() [][]float64 {
	var dataset [][]float64
	file, err := os.Open(s.fileName)
	handleErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	
	rawData, err := reader.ReadAll()
	handleErr(err)

	for _, vals := range rawData{
		dataset = append(dataset,str2flt(vals))
	}
	return dataset
}

func main(){
	snd := sliceNDice{fileName: "data.csv" , splitRatio: 0.75}
	dataset := snd.readCsv()
	fmt.Println(dataset)
}
