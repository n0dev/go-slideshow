package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/n0dev/GoSlideshow/core/picture"
	"github.com/n0dev/GoSlideshow/core/picture/exif"
	"github.com/n0dev/GoSlideshow/logger"
	"github.com/n0dev/GoSlideshow/utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
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
	font       *ttf.Font
	fullscreen bool
}

type imgInfo struct {
	path    string
	H       int32
	W       int32
	texture *sdl.Texture
}

type slideshowInfo struct {
	list    []imgInfo
	current int
}

var slide slideshowInfo

func curImg() *imgInfo {
	return &slide.list[slide.current]
}

func (win *winInfo) addPicture(path string, f os.FileInfo, err error) error {
	if utils.StringInSlice(strings.ToLower(filepath.Ext(path)), validExtensions) {
		slide.list = append(slide.list, imgInfo{path, 0, 0, nil})
	}
	return nil
}

func (win *winInfo) setTitle(position int, total int, path string) {
	win.window.SetTitle(winTitle + " - " + strconv.Itoa(position) + "/" + strconv.Itoa(total) + " - " + filepath.Base(path))
}

func (win *winInfo) setText(message string) {

	black := sdl.Color{A: 0, B: 255, G: 255, R: 255}

	if textSurface, err := win.font.RenderUTF8_Blended(message, black); err == nil {
		texture, err := window.renderer.CreateTextureFromSurface(textSurface)
		if err != nil {
			fmt.Println(err)
		}
		textSurface.Free()

		width := int32(len(message) * 8)
		window.renderer.Copy(texture, nil, &sdl.Rect{X: 2, Y: 2, W: width, H: 16})
		texture.Destroy()
	} else {
		logger.Warning("OMG")
	}
}

func loadImg(win *winInfo, index int) {
	if slide.list[index].texture == nil {
		var err error
		var surface *sdl.Surface

		logger.Trace("load " + slide.list[index].path)

		surface, err = img.Load(slide.list[index].path)
		if err != nil {
			fmt.Printf("Failed to load: %s\n", err)
		}
		defer surface.Free()

		slide.list[index].H = surface.H
		slide.list[index].W = surface.W

		slide.list[index].texture, err = win.renderer.CreateTextureFromSurface(surface)
		if err != nil {
			fmt.Printf("Failed to create texture: %s\n", err)
		}
	}
}

func resetImg(index int) {
	if slide.list[index].texture != nil {
		slide.list[index].texture.Destroy()
		slide.list[index].texture = nil
	}
}

func (win *winInfo) loadAndFreeAround() {
	p1 := utils.Mod(slide.current-1, len(slide.list))
	p2 := utils.Mod(slide.current-1, len(slide.list))
	n1 := utils.Mod(slide.current+1, len(slide.list))
	n2 := utils.Mod(slide.current+2, len(slide.list))
	refresh := utils.IntList{p1, p2, n1, n2}

	// preload the previous and next two images
	for _, idx := range refresh {
		loadImg(win, idx)
	}

	d1 := utils.Mod(slide.current-3, len(slide.list))
	d2 := utils.Mod(slide.current+3, len(slide.list))
	if !refresh.Find(d1) {
		resetImg(d1)
	}
	if !refresh.Find(d2) {
		resetImg(d2)
	}
}

func (win *winInfo) loadCurrentImage() {
	var src, dst sdl.Rect

	// load and display the current image
	loadImg(win, slide.current)

	// Display information of the image
	ww, wh := win.window.GetSize()
	src = sdl.Rect{X: 0, Y: 0, W: curImg().W, H: curImg().H}
	iw, ih := utils.ComputeFitImage(uint32(ww), uint32(wh), uint32(curImg().W), uint32(curImg().H))
	dst = sdl.Rect{X: int32(ww/2 - int(iw)/2), Y: int32(wh/2 - int(ih)/2), W: int32(iw), H: int32(ih)}

	win.renderer.Clear()
	win.renderer.Copy(curImg().texture, &src, &dst)
	win.setText(filepath.Base(curImg().path))
	win.renderer.Present()

	// Preload and free images from the list
	win.loadAndFreeAround()

	/* TEST */
	go func(path string) {
		exif.Open(path)
	}(curImg().path)
}

// Arrange that main.main runs on main thread.
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
	var imagePath, folderPath string

	if fileInfo, _ := os.Stat(inputParam); fileInfo.IsDir() {
		folderPath, _ = filepath.Abs(inputParam)
	} else {
		folderPath, _ = filepath.Abs(filepath.Dir(inputParam))
	}

	// List all pictures in the corresponding folder
	logger.Trace("Listing pictures in " + folderPath)
	filepath.Walk(folderPath, window.addPicture)

	if len(slide.list) != 0 {

		fileInformation, _ := os.Stat(inputParam)
		if fileInformation.IsDir() {
			imagePath, _ = filepath.Abs(slide.list[0].path)
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

	// Load the font library
	if err := ttf.Init(); err != nil {
		logger.Warning("Unable to open font lib")
	}

	window.font, err = ttf.OpenFont("./app/OpenSans-Regular.ttf", 24)
	if err != nil {
		logger.Warning("Unable to load font")
	}

	window.renderer, err = sdl.CreateRenderer(window.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer window.renderer.Destroy()

	// Positioning the current index in list
	for i := range slide.list {
		if slide.list[i].path != imagePath {
			slide.current++
		} else {
			break
		}
	}

	window.setTitle(slide.current+1, len(slide.list), curImg().path)
	window.loadCurrentImage()

	running = true
	for running {
		event = sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_RESIZED || t.Event == sdl.WINDOWEVENT_EXPOSED {
				window.window.SetSize(int(t.Data1), int(t.Data2))

				// Display information of the image
				wWidth, wHeight := window.window.GetSize()

				src = sdl.Rect{X: 0, Y: 0, W: curImg().W, H: curImg().H}
				fitWidth, fitHeight := utils.ComputeFitImage(uint32(wWidth), uint32(wHeight), uint32(curImg().W), uint32(curImg().H))
				dst = sdl.Rect{X: int32(wWidth/2 - int(fitWidth)/2), Y: int32(wHeight/2 - int(fitHeight)/2), W: int32(fitWidth), H: int32(fitHeight)}

				window.renderer.Clear()
				window.renderer.Copy(curImg().texture, &src, &dst)
				window.renderer.Present()
			}

		case *sdl.KeyDownEvent:

			// Get next or previous image
			if t.Repeat == 0 {
				if t.Keysym.Sym == sdl.K_LEFT {
					slide.current = utils.Mod((slide.current - 1), len(slide.list))
				} else if t.Keysym.Sym == sdl.K_RIGHT {
					slide.current = utils.Mod((slide.current + 1), len(slide.list))
				} else if t.Keysym.Sym == sdl.K_PAGEUP {
					if err := picture.RotateImage(curImg().path, picture.CounterClockwise); err != nil {
						logger.Warning(err.Error())
					}
				} else if t.Keysym.Sym == sdl.K_PAGEDOWN {
					if err := picture.RotateImage(curImg().path, picture.Clockwise); err != nil {
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
				} else if t.Keysym.Sym == 105 { // I
					// display image information
				} else if t.Keysym.Sym == sdl.K_ESCAPE {

					if window.fullscreen {
						window.window.SetFullscreen(0)
						window.fullscreen = false
					}

				} else {
					fmt.Printf("%d\n", t.Keysym.Sym)
				}
			}

			window.setTitle(slide.current+1, len(slide.list), curImg().path)
			window.loadCurrentImage()
		}
	}

	return 0
}
