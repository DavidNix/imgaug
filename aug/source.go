package aug

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/pkg/errors"
	"image"
	"os"
	"path/filepath"
)

func EmitSourceImages(dir string) (<-chan image.Image, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, errors.New("no entries in directory")
	}

	ch := make(chan image.Image)
	go func() {
		defer close(ch)
		for _, entry := range entries {
			if entry.IsDir() {
				logInfo("skipping subdirectory %s", entry.Name())
				continue
			}
			img, err := imgio.Open(filepath.Join(dir, entry.Name()))
			if err != nil {
				logErr(errors.Wrap(err, entry.Name()))
				continue
			}
			ch <- img
		}
	}()

	return ch, nil
}
