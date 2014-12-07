package main

import (
	"bufio"
	"os"
	pathpkg "path"
	"path/filepath"
	"strings"

	"github.com/emicklei/hopwatch"
)

// TestCase struct
type TestCase struct {
	// the Test Case name
	Name string
}

// TestSuite struct
type TestSuite struct {
	// the Test Suite name
	Name string
	// file, where Test Suite is defined
	File string
	// Test Cases
	Tests []*TestCase
}

func findMarkdownFiles(folder string) []string {
	var files []string

	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.EqualFold(pathpkg.Ext(info.Name()), ".md") {
				files = append(files, path)
			}
		}
		return nil
	})

	return files
}

func findTestSuites(folder string) []*TestSuite {
	var suites []*TestSuite

	files := findMarkdownFiles(folder)
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		// TODO: replace this code with proper MD parser
		var ts TestSuite
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// get text line
			name := scanner.Text()
			// handle TestSuite
			if strings.HasPrefix(name, "# ") {
				ts = TestSuite{name[2:], path, make([]*TestCase, 0)}
				suites = append(suites, &ts)
			}
			// handle TestCase
			if strings.HasPrefix(name, "## ") {
				if ts.Name != "" {
					t := TestCase{name[3:]}
					ts.Tests = append(ts.Tests, &t)
				}
			}
		}
		hopwatch.Dump(suites).Break()
	}
	return suites
}
