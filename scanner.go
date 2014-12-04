package main

import (
	"bufio"
	"os"
	pathpkg "path"
	"path/filepath"
	"strings"
)

func findMarkdownFiles(folder string) []string {
	var files []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.EqualFold(pathpkg.Ext(info.Name()), "md") {
				files = append(files, path)
			}
		}
		return nil
	})

	return files
}

func listSuites(folder string) []string {
	var suites []string

	files := findMarkdownFiles(folder)
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		name := scanner.Text()
		if strings.HasPrefix(name, "# ") {
			suites = append(suites, name[2:])
		}
	}
	return suites
}
