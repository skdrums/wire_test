package main

import "fmt"

func main() {
	fl := []float64{
		12098,
		512,
		1905,
		162,
		36,
		53,
		7,
	}
	var x float64
	for _, f := range fl {
		x += f
	}
	for _, f := range fl {
		fmt.Println(f / x * 100)
	}

}
