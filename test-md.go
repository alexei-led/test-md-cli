package main

import (
	"fmt"
	"os"

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
	if c.Bool("suites") {
		fmt.Println("Listing all test suites for : ", c.Args().First())
		for _, suite := range listSuites(c.Args().First()) {
			fmt.Println(suite)
		}
	} else {
		fmt.Println("List of all test cases : ", c.Args().First())
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
