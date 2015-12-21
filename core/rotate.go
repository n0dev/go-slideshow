package core

// Rotation test
import (
	"errors"
	"fmt"
	"image"

	"github.com/n0dev/GoSlideshow/logger"
)

// Rotation is
type Rotation int

//
const (
	Clockwise Rotation = iota
	CounterClockwise
)

func rotateClockwise(pic *Picture, rotation Rotation) (image.Image, error) {

	srcBounds := ((*pic).img).Bounds()
	srcX := srcBounds.Dx()
	srcY := srcBounds.Dy()

	switch src := ((*pic).img).(type) {
	case *image.NRGBA:
		fmt.Println("NRGBA")
	case *image.NRGBA64:
		fmt.Println("NRGBA64")
	case *image.RGBA:
		fmt.Println("RGBA")

		dstBounds := image.Rect(srcBounds.Min.Y, srcBounds.Min.X, srcY, srcX)
		dst := image.NewRGBA(dstBounds)

		switch rotation {
		case Clockwise:
			for x := srcBounds.Min.X; x < srcX; x++ {
				for y := srcBounds.Min.Y; y < srcY; y++ {
					dst.Set(srcY-1-y, x, ((*pic).img).At(x, y))
				}
			}

		case CounterClockwise:
			for x := 0; x < srcX; x++ {
				for y := 0; y < srcY; y++ {
					dst.Set(y, srcX-1-x, ((*pic).img).At(x, y))
				}
			}
		}

		return dst, nil

	case *image.RGBA64:
		fmt.Println("RGBA64")
	case *image.Gray:
		fmt.Println("Gray")
	case *image.Gray16:
		fmt.Println("Gray16")
	case *image.YCbCr:
		fmt.Println("YCbCr")
	case *image.Paletted:
		fmt.Println("Paletted")

		dstBounds := image.Rect(srcBounds.Min.Y, srcBounds.Min.X, srcY, srcX)
		dst := image.NewPaletted(dstBounds, src.Palette)

		switch rotation {
		case Clockwise:
			for x := srcBounds.Min.X; x < srcX; x++ {
				for y := srcBounds.Min.Y; y < srcY; y++ {
					dst.Set(srcY-1-y, x, ((*pic).img).At(x, y))
				}
			}

		case CounterClockwise:
			for x := 0; x < srcX; x++ {
				for y := 0; y < srcY; y++ {
					dst.Set(y, srcX-1-x, ((*pic).img).At(x, y))
				}
			}
		}

		return dst, nil

	default:
		fmt.Println("others...")
	}

	return nil, errors.New("error during creation")
}

// RotateImage rotates and save the image at its place
func RotateImage(imagePath string, rotation Rotation) error {

	// Opens and decodes the image given in path or fails
	pic, err := Open(imagePath)
	if err != nil {
		return err
	}

	switch rotation {
	case Clockwise:
		logger.Trace("Rotates " + imagePath + " clockwise")
		if dst, err := rotateClockwise(pic, rotation); err == nil {
			pic.Save(&dst)
		}

	case CounterClockwise:
		logger.Trace("Rotates " + imagePath + " counterclockwise")
		if dst, err := rotateClockwise(pic, rotation); err == nil {
			pic.Save(&dst)
		}

	}

	return nil
}
