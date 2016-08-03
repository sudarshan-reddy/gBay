package gBay

import (
	"encoding/csv"
	//"fmt"
	"os"
	"math/rand"
	//"time"
)

    
type SliceNDice struct{
	FileName string
	SplitRatio float64
	Dataset [][]float64
}

func(s *SliceNDice) ReadCsv() [][]float64 {
	var dataset [][]float64
	file, err := os.Open(s.FileName)
	handleErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	
	rawData, err := reader.ReadAll()
	handleErr(err)

	for _, vals := range rawData{
		dataset = append(dataset,str2flt(vals))
	}
	s.Dataset = dataset
	return dataset
}

func(s SliceNDice) Split() ([][]float64, [][]float64){
	var trainingSet [][]float64
	trainSize := int(float64(len(s.Dataset)) * s.SplitRatio)
	testSet := s.Dataset
	for len(trainingSet) < trainSize{
		index := rand.Intn(len(testSet) - 1)
		trainingSet = append(trainingSet,testSet[index])
		testSet = append(testSet[:index], testSet[index + 1:]...)
	}
	return trainingSet , testSet
}

