deprecated. using https://github.com/bit101/blcairo


# blgg

bitlib library for `gg` (https://github.com/fogleman/gg)

Contains:

* Tons of custom drawing methods beyond `gg` built-ins
* Render library to create:
  * images
  * gifs
  * videos
  * spritesheets
* Utilities for viewing images and videos

Integrates with http://github.com/bit101/bitlib which provides many useful libraries.

Ported from `blgo` (https://github.com/bit101/blgo) which used a fork of `go-cairo` (http://github.com/bit101/go-cairo)

`gg` is a bit simpler than `go-cairo` and has no external dependencies.

## Install

    go get github.com/fogleman/gg
    go get github.com/bit101/blgg
    go get github.com/bit101/bitlib
