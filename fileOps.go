package gBay

import "strconv"

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


