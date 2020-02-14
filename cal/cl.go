package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

var (
	inV   = 1.8
	inA   = 0.02
	span  = 0.01
	evTai = 0.23
)

func main() {
	t := time.Now()
	csvFile, err := os.OpenFile("./cal"+t.String()+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	err = csvFile.Truncate(0)
	writer := csv.NewWriter(csvFile)

	writer.Write([]string{"LEDの効率", "放射束(W)", "光束(lm)", "付加する電力"})

	w := inV * inA
	for i := 1.0; i <= 2.0; i += span {
		line := []string{
			strconv.FormatFloat(i, 'g', 4, 64),
			strconv.FormatFloat(w*i, 'f', 4, 64),
			strconv.FormatFloat(w*i*683.0, 'f', 4, 64),
			strconv.FormatFloat(w-w*i*evTai, 'f', 4, 64),
		}
		writer.Write(line)
	}

	writer.Flush()
}
