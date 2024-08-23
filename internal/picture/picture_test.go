package picture

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	bmpPath  = "../tests/bmp/"
	pngPath  = "../tests/png/"
	gifPath  = "../tests/gif/"
	jpegPath = "../tests/jpeg/"
	tiffPath = "../tests/tiff/"
)

func TestOpenBMP(t *testing.T) {
	files, _ := os.ReadDir(bmpPath)
	for _, f := range files {
		file, _ := filepath.Abs(filepath.Join(bmpPath, f.Name()))
		if _, err := Open(file); err != nil {
			t.Error("cannot open " + file)
		}
	}
}

func TestOpenPNG(t *testing.T) {
	files, _ := os.ReadDir(pngPath)
	for _, f := range files {
		file, _ := filepath.Abs(filepath.Join(pngPath, f.Name()))
		if _, err := Open(file); err != nil {
			t.Error("cannot open " + file)
		}
	}
}

func TestOpenGIF(t *testing.T) {
	files, _ := os.ReadDir(gifPath)
	for _, f := range files {
		file, _ := filepath.Abs(filepath.Join(gifPath, f.Name()))
		if _, err := Open(file); err != nil {
			t.Error("cannot open " + file)
			t.Error(err)
		}
	}
}

func TestOpenJPEG(t *testing.T) {
	files, _ := os.ReadDir(jpegPath)
	for _, f := range files {
		file, _ := filepath.Abs(filepath.Join(jpegPath, f.Name()))
		if _, err := Open(file); err != nil {
			t.Error("cannot open " + file)
		}
	}
}

func TestOpenTIFF(t *testing.T) {
	files, _ := os.ReadDir(tiffPath)
	for _, f := range files {
		file, _ := filepath.Abs(filepath.Join(tiffPath, f.Name()))
		if _, err := Open(file); err != nil {
			t.Error("cannot open " + file)
		}
	}
}
