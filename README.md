# blgg

bitlib library for `gg` (https://github.com/fogleman/gg)

Contains:

* Custom drawing methods beyond `gg` built-ins
* Color library
* Easing library (hello Mr. Penner!)
* Geom library (just points)
* Math library
* Noise library (Perlin and Simplex)
* Random library (PRNG)
* Render and util libraries (create and view images, gifs, videos - requires a few external programs)

Ported from `blgo` (https://github.com/bit101/blgo) which used a fork of `go-cairo` (http://github.com/bit101/go-cairo)

`gg` is a bit simpler than `go-cairo` and has no external dependencies.

## Install

    go get github.com/fogleman/gg
    go get github.com/bit101/blgg
