package main

import (
	"flag"
	"os"

	"github.com/n0dev/GoSlideshow/core"
	"github.com/n0dev/GoSlideshow/logger"
)

// Log usefull information
func init() {

	if name, err := os.Getwd(); err == nil {
		logger.Trace("Execution from " + name)
	}

}

// App starts here
func main() {

	// Parse the command line
	fullScreen := flag.Bool("fullscreen", false, "set in fullscreen mode")
	isSlideshow := flag.Bool("slide", false, "set auto slideshow")
	flag.Parse()

	// Run only if the parameter is a valid file or directory
	args := flag.Args()
	if len(args) == 1 {

		inputParam := args[0]

		// Check wether it is a folder or a file
		if _, err := os.Stat(inputParam); err != nil {
			logger.Error(inputParam + " does not exist")
		} else {
			os.Exit(core.Run(inputParam, *fullScreen, *isSlideshow))
		}

	} else {
		logger.Error("No file or directory given in argument")
	}
	os.Exit(1)
}
