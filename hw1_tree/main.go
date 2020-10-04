package main

import (
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

func dirTree(out io.Writer, rootPath string, needPrintFiles bool) error {
	if needPrintFiles {
		err := printFiles(out, rootPath)
		if err != nil {
			return err
		}
	} else {
		err := printDirectories(out, rootPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func printDirectories(out io.Writer, path string) error {
	//fmt.Println(path)
	open, err := os.Open(path)
	if err != nil {
		return err
	}
	files, err := open.Readdir(-1)

	directories := Filter(files, func(file os.FileInfo) bool {
		return file.IsDir()
	})

	for index, file := range directories {
		if index == len(directories) -1 {
			out.Write([]byte("└───"))
		} else {
			out.Write([]byte("├───"))
		}
		out.Write([]byte(file.Name()))
		out.Write([]byte("\n"))
		printDirectories(out, filepath.Join(path, file.Name()))
	}

	//if e != nil {
	//	return e
	//}
	return nil
}

func printFiles(out io.Writer, path string) error {
	return nil
}

func Filter(files []os.FileInfo, predicate func(os.FileInfo) bool) []os.FileInfo {
	var result []os.FileInfo
	for _, file := range files {
		if predicate(file) {
			result = append(result, file)
		}
	}
	return result
}
