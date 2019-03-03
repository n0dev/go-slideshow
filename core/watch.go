package core

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/n0dev/go-slideshow/logger"
)

// watch is notified on modification on dir path
func watch(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:

				logger.Trace(event.String())

				switch event.Op {
				case fsnotify.Create:
					addPic(event.Name)

				case fsnotify.Rename:
					for i := range slide.list {
						if slide.list[i].path == event.Name {
							//resetImg(i) should be run within the main loop
							slide.list = append(slide.list[:i], slide.list[i+1:]...)
							break
						}
					}

				case fsnotify.Remove:
					for i := range slide.list {
						if slide.list[i].path == event.Name {
							//resetImg(i) should be run within the main loop
							slide.list = append(slide.list[:i], slide.list[i+1:]...)
							break
						}
					}
				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
