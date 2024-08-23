package picture

// Rotation test
import (
	"image"
	"image/draw"
)

// Rotation is
type Rotation int

//
const (
	Clockwise Rotation = iota
	CounterClockwise
)

func rotate(pic *Picture, rotation Rotation) (image.Image, error) {

	// Store useful information
	sb := ((*pic).img).Bounds()
	sw, sh := sb.Bounds().Dx(), sb.Bounds().Dy()
	dw, dh := sh, sw
	b := image.Rect(sb.Min.X, sb.Min.Y, dw, dh)
	dst := image.NewNRGBA(b)

	// Create a new RGBA image to store source
	draw.Draw(dst, b, ((*pic).img), b.Min, draw.Src)

	switch rotation {
	case Clockwise:
		for x := sb.Min.X; x < sw; x++ {
			for y := sb.Min.Y; y < sh; y++ {
				dst.Set(sh-1-y, x, ((*pic).img).At(x, y))
			}
		}

	case CounterClockwise:
		for x := 0; x < sw; x++ {
			for y := 0; y < sh; y++ {
				dst.Set(y, sw-1-x, ((*pic).img).At(x, y))
			}
		}
	}

	return dst, nil
}

// RotateImage rotates and save the image at its place
func RotateImage(imagePath string, rotation Rotation) error {

	// Opens and decodes the image given in path or fails
	pic, err := Open(imagePath)
	if err != nil {
		return err
	}

	dst, err := rotate(pic, rotation)
	if err != nil {
		return err
	}

	return pic.Save(&dst)
}
