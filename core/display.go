package core

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

const (
	winTitle         = "GoSlideshow"
	winDefaultWidth  = 600
	winDefaultHeight = 800
)

// Information about the display window
type winInfo struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	imageList  []string
	fullscreen bool
}

type imgInfo struct {
	path    string
	surface *sdl.Surface
	texture *sdl.Texture
}

var curImage imgInfo

func (win *winInfo) addPicture(path string, f os.FileInfo, err error) error {
	if utils.StringInSlice(strings.ToLower(filepath.Ext(path)), validExtensions) {
		win.imageList = append(win.imageList, path)
	}
	return nil
}

func (win *winInfo) setTitle(position int, total int, path string) {
	win.window.SetTitle(winTitle + " - " + strconv.Itoa(position) + "/" + strconv.Itoa(total) + " - " + filepath.Base(path))
}

func loadImage(window *sdl.Window, renderer *sdl.Renderer, imagePath string) {
	var src, dst sdl.Rect
	var err error

	curImage.surface, err = img.Load(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load: %s\n", err)
	}
	//defer image.Free()

	curImage.texture, err = renderer.CreateTextureFromSurface(curImage.surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
	}
	//defer texture.Destroy()

	// Display information of the image
	wWidth, wHeight := window.GetSize()

	src = sdl.Rect{X: 0, Y: 0, W: curImage.surface.W, H: curImage.surface.H}
	fitWidth, fitHeight := utils.ComputeFitImage(uint32(wWidth), uint32(wHeight), uint32(curImage.surface.W), uint32(curImage.surface.H))
	dst = sdl.Rect{X: int32(wWidth/2 - int(fitWidth)/2), Y: int32(wHeight/2 - int(fitHeight)/2), W: int32(fitWidth), H: int32(fitHeight)}

	renderer.Clear()
	renderer.Copy(curImage.texture, &src, &dst)
	renderer.Present()
}

func init() {
	runtime.LockOSThread()
}

var window winInfo

// Run to launch the app
func Run(inputParam string, fullScreen bool, slideshow bool) int {
	var event sdl.Event
	var running bool
	var src, dst sdl.Rect
	var err error
	var flags uint32 = sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE | sdl.WINDOW_ALLOW_HIGHDPI
	var imagePath string

	folderPath, _ := filepath.Abs(filepath.Dir(inputParam))

	// List all pictures in the corresponding folder
	logger.Trace("Listing pictures in " + folderPath)
	filepath.Walk(folderPath, window.addPicture)

	if len(window.imageList) != 0 {

		fileInformation, _ := os.Stat(inputParam)
		if fileInformation.IsDir() {
			imagePath, _ = filepath.Abs(window.imageList[0])
		} else {
			imagePath, _ = filepath.Abs(inputParam)
		}

	} else {
		logger.Error("No pictures found")
		return 1
	}

	window.window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winDefaultHeight, winDefaultWidth, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.window.Destroy()

	if icon, err := sdl.LoadBMP("./app/icon.bmp"); err == nil {
		window.window.SetIcon(icon)
	}

	window.renderer, err = sdl.CreateRenderer(window.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer window.renderer.Destroy()

	// Positioning the current index in list
	var currentIndex int
	for i := range window.imageList {
		if window.imageList[i] != imagePath {
			currentIndex++
		} else {
			break
		}
	}

	window.setTitle(currentIndex+1, len(window.imageList), window.imageList[currentIndex])
	loadImage(window.window, window.renderer, window.imageList[currentIndex])

	running = true
	for running {
		event = sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_RESIZED {
				window.window.SetSize(int(t.Data1), int(t.Data2))

				// Display information of the image
				wWidth, wHeight := window.window.GetSize()

				src = sdl.Rect{X: 0, Y: 0, W: curImage.surface.W, H: curImage.surface.H}
				fitWidth, fitHeight := utils.ComputeFitImage(uint32(wWidth), uint32(wHeight), uint32(curImage.surface.W), uint32(curImage.surface.H))
				dst = sdl.Rect{X: int32(wWidth/2 - int(fitWidth)/2), Y: int32(wHeight/2 - int(fitHeight)/2), W: int32(fitWidth), H: int32(fitHeight)}

				window.renderer.Clear()
				window.renderer.Copy(curImage.texture, &src, &dst)
				window.renderer.Present()
			}

		case *sdl.KeyDownEvent:

			// Get next or previous image
			if t.Repeat == 0 {
				if t.Keysym.Sym == sdl.K_LEFT {
					currentIndex = utils.Mod((currentIndex - 1), len(window.imageList))
				} else if t.Keysym.Sym == sdl.K_RIGHT {
					currentIndex = utils.Mod((currentIndex + 1), len(window.imageList))
				} else if t.Keysym.Sym == sdl.K_PAGEUP {
					if err := RotateImage(window.imageList[currentIndex], CounterClockwise); err != nil {
						logger.Warning(err.Error())
					}
				} else if t.Keysym.Sym == sdl.K_PAGEDOWN {
					if err := RotateImage(window.imageList[currentIndex], Clockwise); err != nil {
						logger.Warning(err.Error())
					}
				} else if t.Keysym.Sym == 102 { // F

					if window.fullscreen {
						window.window.SetFullscreen(0)
					} else {
						// Go fullscreen
						window.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					}
					window.fullscreen = !window.fullscreen
				} else {
					fmt.Printf("%d\n", t.Keysym.Sym)
				}
			}

			window.setTitle(currentIndex+1, len(window.imageList), window.imageList[currentIndex])
			loadImage(window.window, window.renderer, window.imageList[currentIndex])
		}
	}

	return 0
}
