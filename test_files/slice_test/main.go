package main

import "fmt"

func main() {
	s := []string{}
	s = GetSlice()

	fmt.Println(s)
}

func GetSlice() []string {
	return []string{"1test", "2test", "3test"}
}
