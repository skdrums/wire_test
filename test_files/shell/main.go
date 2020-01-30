package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type FilePath struct {
	Filepath string
}

func main() {
	workDir, _ := filepath.Abs("../annotation_tool_test")

	filePaths := ReadAllFiles(workDir)
	for _, path := range filePaths {
		fmt.Println(path)
	}

	rel := filepath.Base(workDir)
	fmt.Println(rel)
}

func ReadAllFiles(dir string) []string {
	var paths []string
	ob, _ := ioutil.ReadDir(dir)
	for _, o := range ob {
		path := filepath.Join(dir, o.Name())
		if !o.IsDir() {
			paths = append(paths, path)
			continue
		}
		paths = append(paths, ReadAllFiles(path)...)
	}

	return paths
}
