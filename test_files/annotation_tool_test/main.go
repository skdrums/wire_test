package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		working_dir = flag.String("working_dir", "/Users/s07349/go/src/github.com/skdrums/wire_test/test_files/annotation_tool_test/", "working directory")
		csv_path    = flag.String("csv_path", "imagecat", "csv path")
		new_dir     = flag.String("new_dir", "new_image", "")
	)
	flag.Parse()

	file, err := os.Open(*working_dir + *csv_path + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	var csvLine []string
	var info []Info
	for {
		csvLine, err = csvReader.Read()
		if err != nil {
			break
		}
		info = append(info, NewInfo(csvLine))
	}
	fmt.Println(info)

	if err := os.Mkdir(*new_dir, 0777); err != nil {
		fmt.Println(err)
	}

	for n, i := range info {
		if n == 0 {
			continue
		}
		if err := os.Mkdir(*new_dir+"/"+i.Label, 0777); err != nil {
			fmt.Println(err)
		}
		src, err := os.Open(i.ImageDir + i.FileName)
		if err != nil {
			panic(err)
		}
		defer src.Close()

		dst, err := os.Create(*new_dir + "/" + i.Label + "/" + i.FileName)
		if err != nil {
			panic(err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			panic(err)
		}
	}
}

type Info struct {
	ImageDir string
	FileName string
	Label    string
}

func NewInfo(line []string) Info {
	return Info{
		ImageDir: line[0],
		FileName: line[1],
		Label:    line[2],
	}
}
