package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)


//Generic: put these in a separate file later
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
	dataset [][]float64
}

func(s *sliceNDice) readCsv() [][]float64 {
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
	s.dataset = dataset
	return dataset
}

func(s sliceNDice) split() ([][]float64, [][]float64){
	var trainingSet [][]float64
	trainSize := int(float64(len(s.dataset)) * s.splitRatio)
	testSet := s.dataset
	for len(trainingSet) < trainSize{
		index := rand.Intn(len(testSet) - 1)
		trainingSet = append(trainingSet,testSet[index])
		fmt.Println(testSet[index])
		/*delete(testSet , index)*/
	}
	//fmt.Println(trainingSet, testSet)
	return trainingSet , testSet
}

func main(){

	rand.Seed(time.Now().Unix())
	snd := sliceNDice{fileName: "data.csv" , splitRatio: 0.74}
	snd.readCsv()
	snd.split()
	/*fmt.Println(dataset)*/
}
