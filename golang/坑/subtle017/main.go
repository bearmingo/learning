package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	dir := os.Args[1]
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		os.Exit(2)
	}

	var targets []string
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, path)
		return nil
	})

	for _, target := range targets {
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target: ", target, "error: ", err) // error:too many open files
			break
		}

		defer f.Close() // 在每次 for 语句块结束时，不会关闭文件资源

		// 使用 f
	}
}
