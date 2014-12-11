package main

import (
	"fmt"

	"test-md-cli/testmd"

	"github.com/emicklei/hopwatch"

	"github.com/codegangsta/cli"
)

func runTests(c *cli.Context) {
	fmt.Printf("Context : %s\n", c.String("tests"))
}

// print out list of test suites with or w/o test cases
func listTests(c *cli.Context) {
	hopwatch.Display("suites", c.Bool("suites") || c.Bool("s")).Break()
	// list test suites only
	if c.Bool("suites") || c.Bool("s") {
		fmt.Printf("Listing all test suites in : %s\n\n", c.Args().First())
		for _, suite := range testmd.FindTestSuites(c.Args().First()) {
			fmt.Printf(" %s\n", suite.Name)
		}
	} else { // list test suites with all test cases
		fmt.Printf("List of all test cases in : %s\n\n", c.Args().First())
		for _, suite := range testmd.FindTestSuites(c.Args().First()) {
			fmt.Printf("[TS] %s\n", suite.Name)
			for _, test := range suite.Tests {
				fmt.Printf("\t[TC] %s\n", test.Name)
			}
		}
	}
}
