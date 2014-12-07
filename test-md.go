package main

import (
	"fmt"
	"os"

	"github.com/emicklei/hopwatch"

	"github.com/codegangsta/cli"
)

var runCommand = cli.Command{
	Name:      "run",
	ShortName: "r",
	Usage:     "run manual test cases",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "tests, t",
			Value: "*",
			Usage: "test cases to run",
		},
		cli.StringFlag{
			Name:  "out, o",
			Value: ".tmd",
			Usage: "folder to store test case results",
		},
	},
	Action: runTests,
}

var listCommand = cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "list test cases",
	Action:    listTests,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "suites, s",
			Usage: "list only test suites",
		},
	},
}

func runTests(c *cli.Context) {
	fmt.Printf("Context : %s\n", c.String("tests"))
}

func listTests(c *cli.Context) {
	hopwatch.Display("suites", c.Bool("suites") || c.Bool("s")).Break()
	if c.Bool("suites") || c.Bool("s") {
		fmt.Printf("Listing all test suites for : %s\n\n", c.Args().First())
		for _, suite := range findTestSuites(c.Args().First()) {
			fmt.Printf(" %s\n", suite.Name)
		}
	} else {
		fmt.Printf("List of all test cases : %s\n\n", c.Args().First())
		for _, suite := range findTestSuites(c.Args().First()) {
			fmt.Printf("[TS] %s\n", suite.Name)
			for _, test := range suite.Tests {
				fmt.Printf("\t[TC] %s\n", test.Name)
			}
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "test-md"
	app.Version = "0.1"
	app.Usage = "Run test-md manual tests"

	app.Commands = []cli.Command{
		runCommand,
		listCommand,
	}

	splash()
	app.Run(os.Args)
}
