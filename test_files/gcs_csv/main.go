package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

var (
	local_dir = "/Users/s07349/Downloads/CA/CATS/dataset/"
	label     = "non_driver_licence"
	image_uri = "gs://ocr_testdata/"
)

func main() {
	files, err := ioutil.ReadDir(local_dir + label)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, file.Name())
	}

	csvFile, err := os.OpenFile(local_dir+label+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	err = csvFile.Truncate(0)
	writer := csv.NewWriter(csvFile)

	writer.Write([]string{"image_path", "label"})
	for _, path := range paths {
		if strings.Contains(path, ".DS_Store") {
			continue
		}
		line := []string{
			image_uri + label + "/" + path,
			label,
		}
		writer.Write(line)
	}
	writer.Flush()
}
