package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/n0dev/GoSlideshow/logger"
	"github.com/n0dev/GoSlideshow/utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var winTitle = "GoSlideshow"
var winWidth, winHeight int = 800, 600
var imageList []string
var validExtensions = []string{".bmp", ".jpg", ".png", ".gif", ".tif", ".tga"}
var log logger.Logger

func setTitle(window *sdl.Window, position int, total int, path string) {
	window.SetTitle("GoSlideshow - " + strconv.Itoa(position) + "/" + strconv.Itoa(total) + " - " + filepath.Base(path))
}

// Rotation test
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

var image *sdl.Surface
var texture *sdl.Texture

func loadImage(window *sdl.Window, renderer *sdl.Renderer, imagePath string) {
	var src, dst sdl.Rect
	var err error

	image, err = img.Load(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load: %s\n", err)
	}
	//defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
	}
	//defer texture.Destroy()

	// Display information of the image
	wWidth, wHeight := window.GetSize()

	src = sdl.Rect{X: 0, Y: 0, W: image.W, H: image.H}
	fitWidth, fitHeight := utils.ComputeFitImage(uint32(wWidth), uint32(wHeight), uint32(image.W), uint32(image.H))
	dst = sdl.Rect{X: int32(wWidth/2 - int(fitWidth)/2), Y: int32(wHeight/2 - int(fitHeight)/2), W: int32(fitWidth), H: int32(fitHeight)}

	renderer.Clear()
	renderer.Copy(texture, &src, &dst)
	renderer.Present()
}

func init() {
	runtime.LockOSThread()
}

var fullscreen = false

func run(imageName string) int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool
	var err error
	var flags uint32 = sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE | sdl.WINDOW_ALLOW_HIGHDPI

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	// Positioning the current index in list
	var currentIndex int
	for i := range imageList {
		if imageList[i] != imageName {
			currentIndex++
		} else {
			break
		}
	}

	setTitle(window, currentIndex+1, len(imageList), imageList[currentIndex])
	loadImage(window, renderer, imageList[currentIndex])

	running = true
	for running {
		event = sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_RESIZED {
				window.SetSize(int(t.Data1), int(t.Data2))

				var src, dst sdl.Rect

				// Display information of the image
				wWidth, wHeight := window.GetSize()

				src = sdl.Rect{X: 0, Y: 0, W: image.W, H: image.H}
				fitWidth, fitHeight := utils.ComputeFitImage(uint32(wWidth), uint32(wHeight), uint32(image.W), uint32(image.H))
				dst = sdl.Rect{X: int32(wWidth/2 - int(fitWidth)/2), Y: int32(wHeight/2 - int(fitHeight)/2), W: int32(fitWidth), H: int32(fitHeight)}

				renderer.Clear()
				renderer.Copy(texture, &src, &dst)
				renderer.Present()
			}

		case *sdl.KeyDownEvent:

			// Get next or previous image
			if t.Repeat == 0 {
				if t.Keysym.Sym == sdl.K_LEFT {
					currentIndex = utils.Mod((currentIndex - 1), len(imageList))
				} else if t.Keysym.Sym == sdl.K_RIGHT {
					currentIndex = utils.Mod((currentIndex + 1), len(imageList))
				} else if t.Keysym.Sym == sdl.K_PAGEUP {
					rotateImage(imageList[currentIndex], CounterClockwise)
				} else if t.Keysym.Sym == sdl.K_PAGEDOWN {
					rotateImage(imageList[currentIndex], Clockwise)
				} else if t.Keysym.Sym == 102 { // F

					if fullscreen {
						window.SetFullscreen(0)
					} else {
						// Go fullscreen
						window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					}
					fullscreen = !fullscreen
				} else {
					fmt.Printf("%d\n", t.Keysym.Sym)
				}
			}

			setTitle(window, currentIndex+1, len(imageList), imageList[currentIndex])
			loadImage(window, renderer, imageList[currentIndex])
		}
	}

	return 0
}

func visit(path string, f os.FileInfo, err error) error {
	if utils.StringInSlice(strings.ToLower(filepath.Ext(path)), validExtensions) {
		imageList = append(imageList, path)
	}
	return nil
}

func main() {
	log, _ := logger.New("goslideshow.log")

	log.Debug("Starting GoSlideshow")
	if name, err := os.Getwd(); err == nil {
		log.Debug("Execution from " + name)
	}

	args := os.Args
	if len(args) == 2 {
		var imagePath = os.Args[1]

		if _, err := os.Stat(imagePath); err == nil {

			var folderPath, _ = filepath.Abs(filepath.Dir(imagePath))
			log.Debug("Folder path is " + folderPath)
			filepath.Walk(folderPath, visit)

			imageAbsPath, _ := filepath.Abs(imagePath)

			if len(imageList) != 0 {
				os.Exit(run(imageAbsPath))
			} else {
				log.Panic("No pictures found")
			}

		} else {
			log.Panic(imagePath + " does not exist")
		}

	} else {
		log.Panic("No file or directory given in argument")
	}
	os.Exit(1)
}
