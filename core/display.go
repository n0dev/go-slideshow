package core

//#include <stdlib.h>
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"unsafe"

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

var (
	sdlColorWhite = sdl.Color{A: 0, B: 0, G: 0, R: 0}
	sdlColorBlack = sdl.Color{A: 0, B: 255, G: 255, R: 255}
)

// Information about the display window
type winInfo struct {
	window      *sdl.Window
	renderer    *sdl.Renderer
	font        *ttf.Font
	symbols     *ttf.Font
	fullscreen  bool
	displayInfo bool
}

type imgInfo struct {
	path    string
	H       int32
	W       int32
	texture *sdl.Texture
}

// Change the title according to the inputs
func (win *winInfo) setTitle(position int, total int, path string) {
	win.window.SetTitle(winTitle + " - " + strconv.Itoa(position) + "/" + strconv.Itoa(total) + " - " + filepath.Base(path))
}

// Create the texture from text using TTF
func (win *winInfo) renderText(text string) (*sdl.Texture, error) {

	surface, err := win.font.RenderUTF8_Shaded(text, sdlColorBlack, sdlColorWhite)
	defer surface.Free()
	if err != nil {
		return nil, err
	}

	texture, err := win.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

// displayLoading display loading on background
func (win *winInfo) displayLoading() {
	texture, _ := win.renderText("Loading...")
	ww, wh := win.window.GetSize()
	win.renderer.Copy(texture, nil, &sdl.Rect{X: int32(ww/2 - 50), Y: int32(wh/2 - 1), W: 65, H: 20})
	texture.Destroy()
	win.renderer.Present()
}

// displayPictureInfo display all information about the picture
func (win *winInfo) displayPictureInfo() {

	/* Display exif information */
	go func(path string) {
		exif.Open(path)
	}(curImg().path)

	msg := filepath.Base(curImg().path)
	texture, _ := win.renderText(msg)
	width := int32(len(msg) * 8)
	win.renderer.Copy(texture, nil, &sdl.Rect{X: 2, Y: 2, W: width, H: 20})
	texture.Destroy()
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

func (win *winInfo) displayBar() {

	message := "\uF04A \uF04B \uF04E \uF0E2 \uF01E"
	surface, err := win.symbols.RenderUTF8_Shaded(message, sdlColorWhite, sdlColorBlack)

	if err == nil {
		texture, err := win.renderer.CreateTextureFromSurface(surface)
		if err != nil {
			fmt.Println(err)
		}
		surface.Free()

		width := int32(len(message) * 18)
		win.renderer.Copy(texture, nil, &sdl.Rect{X: 120, Y: 500, W: width, H: 40})
		texture.Destroy()
	} else {
		logger.Warning("OMG")
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
	n3 := utils.Mod(slide.current+3, len(slide.list))
	refresh := utils.IntList{p1, p2, n1, n2, n3}

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

func (win *winInfo) loadCurrentImage(render bool) {
	var src, dst sdl.Rect

	// load and display the current image
	loadImg(win, slide.current)

	// Display information of the image
	ww, wh := win.window.GetSize()
	src = sdl.Rect{X: 0, Y: 0, W: curImg().W, H: curImg().H}
	iw, ih := utils.ComputeFitImage(uint32(ww), uint32(wh), uint32(curImg().W), uint32(curImg().H))
	dst = sdl.Rect{X: int32(ww/2 - int(iw)/2), Y: int32(wh/2 - int(ih)/2), W: int32(iw), H: int32(ih)}

	if render {
		win.renderer.Clear()
		win.renderer.Copy(curImg().texture, &src, &dst)
		if window.displayInfo {
			window.displayPictureInfo()
		}
		//window.displayBar()
		win.renderer.Present()
	}

	// Update the window title
	win.setTitle(slide.current+1, len(slide.list), curImg().path)

	// Preload and free images from the list
	win.loadAndFreeAround()
}

// Arrange that main.main runs on main thread.
func init() {

	runtime.LockOSThread()

	// Video only
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		logger.Warning(err.Error())
	}
}

var window winInfo

// MainLoop initializes the SDL package and run the main loop
func MainLoop(fullScreen bool, slideshow bool) int {
	var event sdl.Event
	var src, dst sdl.Rect
	var err error
	var flags uint32 = sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE | sdl.WINDOW_ALLOW_HIGHDPI

	// Load the font library
	if err := ttf.Init(); err != nil {
		logger.Warning("Unable to open font lib")
	}

	window.window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winDefaultHeight, winDefaultWidth, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.window.Destroy()

	// Load resources
	if f, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		icon := filepath.Join(f, "app", "icon.bmp")
		if i, err := sdl.LoadBMP(icon); err == nil {
			window.window.SetIcon(i)
		}

		font := filepath.Join(f, "app", "fonts", "opensans.ttf")
		window.font, err = ttf.OpenFont(font, 14)
		if err != nil {
			logger.Warning("Unable to load " + font)
		}
		window.font.SetKerning(false)

		font = filepath.Join(f, "app", "fonts", "fontawesome.ttf")
		window.symbols, err = ttf.OpenFont(font, 64)
		if err != nil {
			logger.Warning("Unable to load " + font)
		}
	}

	window.renderer, err = sdl.CreateRenderer(window.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer window.renderer.Destroy()

	window.displayInfo = false
	window.displayLoading()
	window.setTitle(slide.current+1, len(slide.list), curImg().path)
	window.loadCurrentImage(false)

	// Declare if the image needs to be updated
	var update = false
	var running = true

	for running {
		event = sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false

		case *sdl.DropEvent:
			p := (*C.char)(t.File)
			s := C.GoString(p)
			C.free(unsafe.Pointer(p))

			// Check if picture already in list
			found := false
			for i := range slide.list {
				if slide.list[i].path == s {
					found = true
					slide.current = i
					update = true
					break
				}
			}
			if !found {
				if err := addPic(s); err != nil {
					sdl.ShowSimpleMessageBox(sdl.MESSAGEBOX_INFORMATION, "File dropped on window", "Cannot add "+s, window.window)
				} else {
					slide.current = len(slide.list) - 1
					update = true
				}
			}

		/*case *sdl.MouseMotionEvent:
			fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)

		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)*/

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

				if window.displayInfo {
					window.displayPictureInfo()
					window.renderer.Present()
				}
			}

		case *sdl.KeyDownEvent:

			// Get next or previous image
			if t.Repeat == 0 {
				if t.Keysym.Sym == sdl.K_LEFT {
					slide.current = utils.Mod((slide.current - 1), len(slide.list))
					update = true
				} else if t.Keysym.Sym == sdl.K_RIGHT {
					slide.current = utils.Mod((slide.current + 1), len(slide.list))
					update = true
				} else if t.Keysym.Sym == sdl.K_PAGEUP {
					if err := picture.RotateImage(curImg().path, picture.CounterClockwise); err != nil {
						logger.Warning(err.Error())
					} else {
						resetImg(slide.current)
					}
					update = true
				} else if t.Keysym.Sym == sdl.K_PAGEDOWN {
					if err := picture.RotateImage(curImg().path, picture.Clockwise); err != nil {
						logger.Warning(err.Error())
					} else {
						resetImg(slide.current)
					}
					update = true
				} else if t.Keysym.Sym == 102 { // F

					if window.fullscreen {
						window.window.SetFullscreen(0)
					} else {
						// Go fullscreen
						window.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					}
					window.fullscreen = !window.fullscreen
				} else if t.Keysym.Sym == 105 { // I

					window.displayInfo = !window.displayInfo
					if window.displayInfo {
						fmt.Println("Toggle info: on")
						window.displayPictureInfo()
						window.renderer.Present()
					} else {
						fmt.Println("Toggle info: off")
						update = true
					}

				} else if t.Keysym.Sym == sdl.K_ESCAPE {

					if window.fullscreen {
						window.window.SetFullscreen(0)
						window.fullscreen = false
					}

				} else {
					fmt.Printf("%d\n", t.Keysym.Sym)
				}
			}
		}
		if update {
			window.loadCurrentImage(true)
			update = false
		}
	}

	return 0
}
