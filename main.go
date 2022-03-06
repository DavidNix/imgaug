package main

import (
	"flag"
	"fmt"
)

const usage = `Augment images from target directory and write augmented results to a new directory.

Usage:
  imgaug [target_dir] [flags]

Flags:
  --help    show this help message
`

type config struct {
	Help bool
}

func main() {
	var cfg config
	{
		flag.BoolVar(&cfg.Help, "help", false, "print help")

		flag.Parse()
	}

	switch {
	case cfg.Help:
		fmt.Printf(usage)
	}
}
