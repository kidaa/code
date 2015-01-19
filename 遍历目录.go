package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunc(path string, info os.FileInfo, err error) error {
	return nil
}

type Walker struct {
	directories []string
	files       []string
}

func main() {
	flag.Parse()
	//    root := flag.Arg(0)
	walker := new(Walker)
	path := "."
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			walker.directories = append(walker.directories, path)
		} else {
			walker.files = append(walker.files, path)
		}
		return nil
	})
	fmt.Printf("found %d dir and  %d files\n", len(walker.directories), len(walker.files))
	for i := 0; i < len(walker.files); i += 1 {
		fmt.Println(walker.files[i])
	}
}
