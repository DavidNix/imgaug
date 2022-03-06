package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/DavidNix/imgaug/aug"
	"os"
	"path/filepath"
	"time"
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
	SourceDir string
}

func (cfg config) Validate() error {
	if len(cfg.SourceDir) == 0 {
		return errors.New("missing source directory, --dir")
	}
	return nil
}

func exit(err error) {
	fmt.Printf("ERROR: %v\n\n", err)
	fmt.Printf(usage)
	os.Exit(1)
}

func main() {
	var cfg config

	{
		flag.StringVar(&cfg.SourceDir, "dir", "", "target directory containing images")
		flag.BoolVar(&cfg.Help, "help", false, "print help")

		flag.Parse()
	}

	if cfg.Help {
		fmt.Printf(usage)
		return
	}

	if err := cfg.Validate(); err != nil {
		exit(err)
	}

	destDir, err := validateDirectories(cfg)
	if err != nil {
		exit(err)
	}

	srcCh, err := aug.EmitSourceImages(cfg.SourceDir)
	if err != nil {
		exit(err)
	}

	for range srcCh {
		fmt.Println("got image!")
	}

	fmt.Println("augmented images saved to", destDir)
}

func validateDirectories(cfg config) (string, error) {
	_, err := os.Stat(cfg.SourceDir)
	if err != nil {
		return "", err
	}

	dest := filepath.Join(cfg.SourceDir, fmt.Sprintf("augmented-%d", time.Now().Unix()))
	err = os.MkdirAll(dest, 0666)
	if err != nil {
		return "", err
	}
	return dest, nil
}
