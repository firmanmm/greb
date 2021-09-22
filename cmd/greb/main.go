package main

import (
	"fmt"
	"log"
	"os"

	"github.com/firmanmm/greb"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Generate HTTP Request Binding"
	app.Version = "AA (Always Alpha)"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "out",
			Usage: "Set output file, eg : --out=ok.go",
		},
		&cli.StringFlag{
			Name:     "in",
			Usage:    "Set input file, eg : --in=random.gerb",
			Required: true,
		},
	}
	app.Action = func(c *cli.Context) error {
		inFile := c.String("in")
		res, err := greb.Generate(inFile)
		if err != nil {
			return err
		}
		outFile := c.String("out")
		if outFile == "" {
			fmt.Println(res)
		} else {
			return os.WriteFile(outFile, []byte(res), 0666)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err.Error())
	}
}
