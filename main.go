package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	debugFlag = false
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1"
	app.Usage = "tool for managing .mcm character sets for MAX7456"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "force, f",
			Usage:       "Overwrite output files without asking",
			Destination: &forceFlag,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Print debug messages",
			Destination: &debugFlag,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "extract",
			Usage:     "Extract all characters to individual images",
			ArgsUsage: "<input.mcm> <output-dir>",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "add-blanks, b",
					Usage: "Include blank characters in the extracted files",
				},
			},
			Action: extractAction,
		},
		{
			Name:      "build",
			Usage:     "Build a .mcm from the files in the given directory",
			ArgsUsage: "<input-dir> <output.mcm>",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-blanks, b",
					Usage: "Don't fill missing characters with blanks",
				},
			},
			Action: buildAction,
		},
		{
			Name:      "png",
			Usage:     "Generate a .png from an .mcm",
			ArgsUsage: "<input.mcm> <output.png>",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "margin, m",
					Value: 1,
					Usage: "Margin between each character",
				},
				cli.IntFlag{
					Name:  "columns, c",
					Value: 16,
					Usage: "Number of columns in the output image",
				},
			},
			Action: pngAction,
		},
	}
	app.Run(os.Args)
}
