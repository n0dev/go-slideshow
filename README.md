# GoSlideshow

![Build Status](https://travis-ci.org/n0dev/GoSlideshow.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/n0dev/GoSlideshow/badge.svg?branch=master&service=github)](https://coveralls.io/github/n0dev/GoSlideshow?branch=master)

Simple cross-platform (just tested on Linux yet) slideshow for pictures written in go.

**This is not production ready yet.**

## Install on Linux

To install just follow the steps. Builds with go 1.12:

```bash
apt-get install libsdl2{,-mixer,-image,-ttf}-dev
go get github.com/n0dev/go-slideshow
```

## Usage

Just run `goslideshow` indicating a file or a folder in parameter.

```text
| Keys          | Are                           |
| ------------- | ----------------------------- |
| `🠊`           | Next picture                  |
| `🠈`           | Previous picture              |
| `f`           | Toggle fullscreen             |
| `i`           | Toggle information on picture |
```

## License

Copyright (c) 2015 Nicolas Hess nicolas.hess@gmail.com goslideshow is released under a BSD style license.
