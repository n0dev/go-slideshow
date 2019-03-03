package core

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/n0dev/go-slideshow/utils"
)

var validExtensions = []string{
	".bmp",
	".jpg",
	".jpeg",
	".png",
	".gif",
	".tif",
	".tga",
}

// Various errors returned by Scan
var (
	ErrSlideEmpty    = errors.New("no picture has been found")
	ErrNotCompatible = errors.New("picture in parameter is not compatible")
)

type slideshowInfo struct {
	list    []imgInfo // array of image info
	current int       // current index in the slideshow
}

var slide slideshowInfo

// Alias to return the image currently displayed
func curImg() *imgInfo {
	return &slide.list[slide.current]
}

// addPic checks if the picture has the right extension and adds it
func addPic(p string) error {
	if !utils.StringInSlice(strings.ToLower(filepath.Ext(p)), validExtensions) {
		return ErrNotCompatible
	}
	a, _ := filepath.Abs(p)
	slide.list = append(slide.list, imgInfo{a, 0, 0, nil})
	return nil
}

// addDir scans and add dir for pictures to slide
func addDir(p string, recurse *bool) {

	if *recurse {
		filepath.Walk(p, func(img string, f os.FileInfo, err error) error {
			addPic(img)
			return nil
		})

	} else {
		if files, err := ioutil.ReadDir(p); err == nil {
			for _, img := range files {
				addPic(filepath.Join(p, img.Name()))
			}
		}
	}

	go watch(p)
}

// Scan loads the pictures' list here or returns an error
func Scan(p []string, recurse *bool) error {

	if len(p) == 0 {
		return ErrSlideEmpty

	} else if len(p) == 1 {

		var input = p[0]
		var baseDir string

		// Check if the parameter exists or is a file or directory
		if i, err := os.Stat(input); err != nil {
			return err

		} else if i.IsDir() {
			baseDir, _ = filepath.Abs(filepath.Clean(input))
			addDir(baseDir, recurse)

			// Quick check if picture has been found
			if len(slide.list) != 0 {
				slide.current = 0
			} else {
				return ErrSlideEmpty
			}

		} else {
			fileAbs, _ := filepath.Abs(filepath.Clean(input))
			baseDir, _ = filepath.Abs(filepath.Dir(fileAbs))
			addDir(baseDir, recurse)

			// Quick check if picture has been found
			if len(slide.list) != 0 {
				found := false
				for i := range slide.list {
					if slide.list[i].path == fileAbs {
						slide.current = i
						found = true
						break
					}
				}
				if !found {
					return ErrNotCompatible
				}
			} else {
				return ErrSlideEmpty
			}
		}

	} else if len(p) > 1 {

		// Check if each parameter exists or is a file or directory
		for _, f := range p {
			if i, err := os.Stat(f); err != nil {
				return err

			} else if i.IsDir() {
				baseDir, _ := filepath.Abs(f)
				addDir(baseDir, recurse)

			} else {
				addPic(f)
			}
		}

		// Quick check if picture has been found
		if len(slide.list) != 0 {
			slide.current = 0
		} else {
			return ErrSlideEmpty
		}
	}
	return nil
}
