package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/n0dev/go-slideshow/core"
	"github.com/n0dev/go-slideshow/logger"
)

var fullScreen *bool
var isSlideshow *bool
var recurse *bool

// Starts the logger and parse the command line
func init() {

	if name, err := os.Getwd(); err == nil {
		logger.Trace("Execution from " + name)
	}

	fullScreen = flag.Bool("f", false, "set in fullscreen mode")
	isSlideshow = flag.Bool("s", false, "set auto slideshow")
	recurse = flag.Bool("r", false, "scan pictures recursively in folders")
	flag.Parse()
}

// App starts here
func main() {

	// Folder and parameter scanning
	if err := core.Scan(flag.Args(), recurse); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Starts the main loop
	core.MainLoop(*fullScreen, *isSlideshow)
}
