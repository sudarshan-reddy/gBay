package gBay

import (
	"encoding/csv"
	"os"
	"math/rand"
	"time"
	//"fmt"
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
	rand.Seed(time.Now().Unix())
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

func(s SliceNDice) ByClass(dataset [][]float64) map[int][][]float64 {
	clsSep := make(map[int][][]float64)
	clsSum := make(map[int][][]float64)
	for  _ , v := range dataset{
		key := int(v[len(v) - 1])
		clsSep[key] = append(clsSep[key], v)
	}
	
	for k,  v := range clsSep{
		clsSum[k] = Summarize(v)		
	}
	return clsSum
}



/*Usage doodle- to be removed

func main(){
	snd := gBay.SliceNDice{FileName: "data.csv" , SplitRatio: 0.74}
	snd.ReadCsv()
	a, _ := snd.Split()
	fmt.Println(snd.ByClass(a))

}

*/
