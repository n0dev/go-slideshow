# GoSlideshow

Simple cross-platform (just tested on Linux yet) slideshow written in go.

**This is not ready for use yet.**

### Install on Linux

To install just follow the steps:

`apt-get install libsdl2{,-mixer,-image,-ttf}-dev`

`go get -v github.com/veandco/go-sdl2/sdl{,_mixer,_image,_ttf}`

`go get github.com/n0dev/GoSlideshow`

### Usage

Just run `goslideshow` indicating a file or a folder in parameter.

### Todo

- [ ] Properly refacto the code sources
- [ ] Add tests & coverage
- [ ] Build with Travis CI
- [ ] Platform independent
- [ ] Re-sizable window
- [x] Image should have its size otherwise size of screen
- [ ] Add rotate function with auto-save
- [ ] Auto slideshow for a directory with configurable timer
- [ ] Folder tree
