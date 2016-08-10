#H2 Naive Bayes classifier implentation in Go 

Takes in a dataset/CSV file, (resembling a numpy matrix) and returns predictions
and accuracy.

TODO: 

Only used Gaussian PDF. Implement bernouli and multinomial as well (mathOPs.go change)


Fix floating point underflow


Write an email spam example and link it here

Installation Instructions

```sh
go get https://github.com/sudsred/gBay
```

Usage Example: 

```go
package main

import ( "github/sudsred/gBay"
         "fmt"
         )

func main(){
	fmt.Println("Test")
	//data.csv is a sample file. It is attached with this repo as
	//a template example. Any similar file should work fine
	snd := gBay.SliceNDice{FileName: "data.csv" , SplitRatio: 0.75}
	snd.ReadCsv()
	training , test := snd.Split()
	//prepare model
	summaries := snd.ByClass(training)
	//predictions	
	predictions := gBay.GetPredictions(summaries, test)
	fmt.Println(gBay.GetAccuracy(test, predictions))
}
```

Questions/Suggestions?

Email: sudar.theone@gmail.com
