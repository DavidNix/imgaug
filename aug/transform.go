package aug

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"image"
	"path/filepath"
)

type Transformer struct {
	Dir   string
	total int64
}

func (tf *Transformer) Augment(in <-chan image.Image) int64 {
	for orig := range in {
		for _, img := range tf.baseImages(orig) {
			tf.savePNG(img)
		}
	}
	return tf.total
}

func (tf *Transformer) baseImages(orig image.Image) (images []image.Image) {
	for _, img := range []image.Image{orig, transform.FlipH(orig)} {
		images = append(
			images,
			img,
			transform.Rotate(img, -10, nil),
			transform.Rotate(img, -20, nil),
			transform.Rotate(img, -30, nil),
			transform.Rotate(img, 10, nil),
			transform.Rotate(img, 20, nil),
			transform.Rotate(img, 30, nil),
		)
	}
	return images
}

func (tf *Transformer) savePNG(img image.Image) {
	fname := filepath.Join(tf.Dir, uuid.NewString()+".png")
	if err := imgio.Save(fname, img, imgio.PNGEncoder()); err != nil {
		logErr(errors.Wrap(err, "image save"))
		return
	}
	tf.total++
}
