package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	date     = "1995/06/18"
	testYear = "H01.12.01"
	test     = 966
	byo      = "10.8秒"
	byo1     = "10.8"

	adMap = map[string]int64{
		"T": 1911,
		"S": 1925,
		"H": 1988,
		"R": 2018,
	}
)

//func main() {
//	year := date[0:4]
//	month := date[5:7]
//	day := date[8:10]
//	fmt.Println(year + month + day)
//}

func main() {
	wareki := testYear[0:1]
	var ad int64
	for key, value := range adMap {
		if key == wareki {
			ad = value
			break
		}
	}

	year, err := strconv.ParseInt(testYear[1:3], 10, 64)
	if err != nil {
		panic(err)
	}

	month := testYear[4:6]
	day := testYear[7:9]

	fmt.Println(fmt.Sprintf("%4d%s%s", ad+year, month, day))
	fmt.Println(fmt.Sprintf("%04d", test))
	fmt.Println(len(date))

	testint := int64(11)
	fmt.Println(string(testint))

	is := (testint > 10)
	fmt.Println(is)

	now := time.Now().Format("20060102")
	fmt.Println(now)
}
