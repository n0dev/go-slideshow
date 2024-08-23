package utils

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Mod is the function which always return the positive modulus
func Mod(a int, b int) int {
	rem := a % b
	if rem < 0 {
		rem += b
	}
	return rem
}

// GetAppData returns the application folder for this application. Creates it
// otherwise. Can return an error if the path has no write rights.
func GetAppData() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the path
	f := filepath.Join(usr.HomeDir, ".goslideshow")
	_, err = os.Stat(f)
	if os.IsNotExist(err) {
		if err := os.Mkdir(f, 0711); err != nil {
			return "", err
		}
	}
	return f, nil
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
