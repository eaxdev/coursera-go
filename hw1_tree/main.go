package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, needPrintFiles bool) error {

	if needPrintFiles {
		err := printFiles(out, path)
		if err != nil {
			return err
		}
	} else {
		err := printDirectories(out, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func printDirectories(out io.Writer, path string) error {
	e := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		return nil
	})
	if e != nil {
		return e
	}
	return nil
}

func printFiles(out io.Writer, path string) error {
	return nil
}