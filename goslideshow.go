package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/n0dev/GoSlideshow/utils"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle = "GoSlideshow"
var winWidth, winHeight int = 800, 600
var imageList []string

func run(imageName string) int {
	runtime.LockOSThread()
	var window *sdl.Window
	var renderer *sdl.Renderer
	var image *sdl.Surface
	var texture *sdl.Texture
	var src, dst sdl.Rect
	var event sdl.Event
	var running bool
	var err error

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
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

	image, err = sdl.LoadBMP(imageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
		return 3
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	src = sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
	dst = sdl.Rect{X: 100, Y: 50, W: 512, H: 512}

	renderer.Clear()
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	// Positioning the current index in list
	var currentIndex int

	for i := range imageList {
		if imageList[i] != imageName {
			currentIndex++
		} else {
			break
		}
	}

	window.SetTitle(winTitle + " - " + filepath.Base(imageList[currentIndex]))

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			/*case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)*/
			case *sdl.KeyDownEvent:
				/*fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)*/

				if t.Repeat == 0 {
					if t.Keysym.Sym == sdl.K_LEFT {
						currentIndex = utils.Mod((currentIndex - 1), len(imageList))
					} else if t.Keysym.Sym == sdl.K_RIGHT {
						currentIndex = utils.Mod((currentIndex + 1), len(imageList))
					}
				}

				//fmt.Printf("index: %d\timage:%s\n", currentIndex, imageList[currentIndex])

				window.SetTitle(winTitle + " - " + filepath.Base(imageList[currentIndex]))

				image, err = sdl.LoadBMP(imageList[currentIndex])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
					return 3
				}
				defer image.Free()

				texture, err = renderer.CreateTextureFromSurface(image)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
					return 4
				}
				defer texture.Destroy()

				src = sdl.Rect{X: 0, Y: 0, W: 512, H: 512}
				dst = sdl.Rect{X: 100, Y: 50, W: 512, H: 512}

				renderer.Clear()
				renderer.Copy(texture, &src, &dst)
				renderer.Present()
			}
		}
	}

	return 0
}

func visit(path string, f os.FileInfo, err error) error {
	if filepath.Ext(path) == ".bmp" {
		fmt.Printf("Add: %s\n", path)
		imageList = append(imageList, path)
	}
	return nil
}

func main() {
	fmt.Println("[+] Starting GoSlideshow")
	if name, err := os.Getwd(); err == nil {
		fmt.Println("[+] Execution from " + name)
	}
	args := os.Args
	if len(args) == 2 {
		var imagePath = os.Args[1]

		if _, err := os.Stat(imagePath); err == nil {

			var folderPath, _ = filepath.Abs(filepath.Dir(imagePath))
			fmt.Println("[+] Folder path is " + folderPath)
			filepath.Walk(folderPath, visit)

			imageAbsPath, _ := filepath.Abs(imagePath)
			os.Exit(run(imageAbsPath))
		} else {
			fmt.Println("[*] " + imagePath + " does not exist")
		}

	} else {
		fmt.Println("[*] No file or directory given in argument")
	}
	os.Exit(1)
}
