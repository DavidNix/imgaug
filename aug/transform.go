package aug

import (
	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/effect"
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
			for _, v := range []float64{-0.3, -0.15, 0.15, 0.3} {
				var (
					bright   = adjust.Brightness(img, v)
					contrast = adjust.Contrast(img, v)
				)
				tf.savePNG(bright)
				tf.savePNG(contrast)
				tf.savePNG(adjust.Brightness(contrast, v))
				tf.savePNG(adjust.Contrast(bright, v))
			}
		}
	}
	return tf.total
}

func (tf *Transformer) baseImages(orig image.Image) (images []image.Image) {
	for _, img := range []image.Image{effect.Sharpen(orig), effect.Sharpen(transform.FlipH(orig))} {
		images = append(
			images,
			img,
			transform.Rotate(img, -10, nil),
			transform.Rotate(img, -20, nil),
			transform.Rotate(img, 10, nil),
			transform.Rotate(img, 20, nil),
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

func (tf *Transformer) adjustBrightness(img image.Image) {
	for _, v := range []float64{-0.3, -0.15, 0.15, 0.3} {
		tf.savePNG(adjust.Brightness(img, v))
	}
}

func (tf *Transformer) adjustContrast(img image.Image) {
	for _, v := range []float64{-0.3, -0.15, 0.15, 0.3} {
		tf.savePNG(adjust.Contrast(img, v))
	}
}
