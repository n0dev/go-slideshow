# Go Slideshow

Simple cross-platform slideshow to display your pictures in go.

**This is not production ready yet.**

## Setup

First install the dependencies for your system:

```bash
apt-get install libsdl2{,-mixer,-image,-ttf}-dev
```

Or MacOS

```bash
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
```

## Build and run

```bash
make build
./bin/go-slideshow <file or folder>
```

## Usage

```text
| Keys       | Com                                 |
| ---------- | ----------------------------------- |
| →          | Next picture                        |
| ←          | Previous picture                    |
| f          | Toggle fullscreen                   |
| i          | Toggle information on picture       |
```