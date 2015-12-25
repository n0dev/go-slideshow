package utils

// Mod is the function which always return the positive modulus
func Mod(a int, b int) int {
	rem := a % b
	if rem < 0 {
		rem += b
	}
	return rem
}

// StringInSlice returns true if the string a is in slice, else false
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntList for an integer list
type IntList []int

// Find returns true if the integer is in list
func (l *IntList) Find(i int) bool {
	for _, a := range *l {
		if a == i {
			return true
		}
	}
	return false
}

// ComputeFitImage returns Width and Height for the image to fit in the window, keeping the actual ratio
func ComputeFitImage(winWidth, winHeight, imageWidth, imageHeight uint32) (uint32, uint32) {

	if winHeight == 0 || imageHeight == 0 {
		return 0, 0
	}

	// If the image already fits, returns its dimension
	if imageWidth < winWidth && imageHeight < winHeight {
		return imageWidth, imageHeight
	}

	var winRatio = float64(winWidth) / float64(winHeight)
	var imageRatio = float64(imageWidth) / float64(imageHeight)

	if winRatio > imageRatio { // Will be limited in height, else in width
		return uint32(float64(winHeight) * imageRatio), winHeight
	}

	return winWidth, uint32(float64(winWidth) / imageRatio)
}
