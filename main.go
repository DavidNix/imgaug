package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const usage = `Augment images from source directory and write augmented results to a new directory.

Usage:
  imgaug [flags]

Flags:
  --dir     (required) source directory containing images
  --help    show this help message
`

type config struct {
	Help      bool
	TargetDir string
}

func (cfg config) Validate() error {
	if len(cfg.TargetDir) == 0 {
		return errors.New("missing source directory, --dir")
	}
	return nil
}

func printErr(err error) {
	fmt.Printf("ERROR: %v\n\n", err)
	fmt.Printf(usage)
	os.Exit(1)
}

func main() {
	var cfg config

	{
		flag.StringVar(&cfg.TargetDir, "dir", "", "target directory containing images")
		flag.BoolVar(&cfg.Help, "help", false, "print help")

		flag.Parse()
	}

	if cfg.Help {
		fmt.Printf(usage)
		return
	}

	if err := cfg.Validate(); err != nil {
		printErr(err)
	}
}
