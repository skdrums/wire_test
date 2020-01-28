package main

import (
	"encoding/csv"
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var (
		image_dir = flag.String("image_dir", "./", "image_directory including images")
		label     = flag.String("label", "no_label", "label")
	)
	flag.Parse()
	files, err := ioutil.ReadDir(*image_dir)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, file.Name())
	}

	csvFile, err := os.OpenFile(*image_dir+*label+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	err = csvFile.Truncate(0)
	writer := csv.NewWriter(csvFile)

	writer.Write([]string{"image_dir", "image_name", "label"})
	for _, path := range paths {
		if strings.Contains(path, ".DS_Store") {
			continue
		}
		line := []string{
			*image_dir + "/",
			path,
			*label,
		}
		writer.Write(line)
	}
	writer.Flush()
}
