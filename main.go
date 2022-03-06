package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/DavidNix/imgaug/aug"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
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

	destDir, err := buildDestination(cfg)
	if err != nil {
		exit(err)
	}

	ch, err := aug.EmitSourceImages(cfg.SourceDir)
	if err != nil {
		exit(err)
	}

	var (
		total int64
		wg    sync.WaitGroup
		cpus  = runtime.NumCPU()
	)

	wg.Add(cpus)
	for i := 0; i < cpus; i++ {
		go func() {
			subtotal := (&aug.Transformer{Dir: destDir}).Augment(ch)
			atomic.AddInt64(&total, subtotal)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("%d augmented images saved to %s\n", total, destDir)
}

func buildDestination(cfg config) (string, error) {
	_, err := os.Stat(cfg.SourceDir)
	if err != nil {
		return "", err
	}

	dest := filepath.Join(cfg.SourceDir, fmt.Sprintf("augmented-%d", time.Now().Unix()))
	err = os.MkdirAll(dest, 0776)
	if err != nil {
		return "", err
	}
	return dest, nil
}
