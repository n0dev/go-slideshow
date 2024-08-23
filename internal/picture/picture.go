package picture

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// Format f
type Format int

// JPEG
const (
	JPEG Format = iota
	PNG
	GIF
	TIFF
	BMP
)

// Picture the struct for the picture
type Picture struct {
	path   string
	format Format
	img    image.Image
}

var formats = map[string]Format{
	"jpeg": JPEG,
	"png":  PNG,
	"tiff": TIFF,
	"bmp":  BMP,
	"gif":  GIF,
}

// Open returns a picture if it decodes it
func Open(imagePath string) (*Picture, error) {

	// Opens the image given in path or fails
	fImg, err := os.Open(imagePath)
	defer fImg.Close()
	if err != nil {
		return nil, err
	}

	// Decodes the images, rotates if succedded
	src, format, err := image.Decode(fImg)
	if err == nil {

		f, ok := formats[format]
		if !ok {
			return nil, errors.New("format " + format + " is not supported")
		}

		return &Picture{path: imagePath, format: f, img: src}, nil
	}

	return nil, err
}

// Save the image given
func (pic *Picture) Save(img *image.Image) error {

	// Create desination
	f, err := os.Create(pic.path)
	defer f.Close()
	if err != nil {
		return err
	}

	switch pic.format {
	case JPEG:
		return jpeg.Encode(f, *img, &jpeg.Options{Quality: 95})
	case PNG:
		return png.Encode(f, *img)
	case GIF:
		return gif.Encode(f, *img, &gif.Options{NumColors: 256})
	case TIFF:
		return tiff.Encode(f, *img, &tiff.Options{Compression: tiff.Deflate, Predictor: true})
	case BMP:
		return bmp.Encode(f, *img)
	}

	return errors.New("unsupported format")
}
