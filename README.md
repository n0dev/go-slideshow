# GoSlideshow

![Build Status](https://travis-ci.org/n0dev/GoSlideshow.svg?branch=master)

Simple cross-platform (just tested on Linux yet) slideshow written in go.

**This is not ready for use yet.**

### Install on Linux

To install just follow the steps:

`apt-get install libsdl2{,-mixer,-image,-ttf}-dev`

`go get -v github.com/veandco/go-sdl2/sdl{,_mixer,_image,_ttf}`

`go get github.com/n0dev/GoSlideshow`

### Usage

Just run `goslideshow` indicating a file or a folder in parameter.


| Keys          | Are                |
| ------------- | ------------------ |
| `🠊`           | Next picture       |
| `🠈`           | Previous picture   |
| `f`           | Toggle fullscreen  |

### License

Copyright (c) 2015 Nicolas Hess nicolas.hess@gmail.com goslideshow is released under a BSD style license.
