package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	minV = 1.3
	maxV = 2.0
	span = 0.05
)

func main() {
	t := time.Now()
	csvFile, err := os.OpenFile("./ep2_pm"+t.String()+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	err = csvFile.Truncate(0)
	writer := csv.NewWriter(csvFile)

	writer.Write([]string{"電圧", "電流", "電力"})

	var v, a, w string
	for i := minV; i <= maxV; i += span {
		fmt.Println("電圧:", i, "\n電流は？")
		fmt.Scan(&a)
		if a == "quit" {
			break
		}
		fmt.Println("電力は？")
		fmt.Scan(&w)
		v = strconv.FormatFloat(i, 'f', 2, 64)
		line := []string{
			v,
			a,
			w,
		}
		if w == "quit" {
			break
		}
		writer.Write(line)
	}

	writer.Flush()
}
