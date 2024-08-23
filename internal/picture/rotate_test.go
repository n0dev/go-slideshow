package picture

import "testing"

func TestRotateImage(t *testing.T) {
	RotateImage("../app/donald.gif", Clockwise)
}
