package core

// Rotation test
import "fmt"

// Rotation is
type Rotation int

//
const (
	Clockwise Rotation = iota
	CounterClockwise
)

func rotateImage(imagePath string, rotation Rotation) {
	switch rotation {
	case Clockwise:
		fmt.Println("Clockwise rotation")
	case CounterClockwise:
		fmt.Println("CounterClockwise rotation")
	}
}
