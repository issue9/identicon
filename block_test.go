// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"testing"

	"github.com/issue9/assert"
)

func drawImage(file string, fn blockFunc, a *assert.Assertion) {
	p := []color.Color{color.RGBA{255, 0, 0, 100}, color.RGBA{0, 255, 255, 100}}
	img := image.NewPaletted(image.Rect(0, 0, 128*4, 128), p) // 横向4张图片大小
	a.NotNil(img)

	for i := 0; i < 4; i++ {
		fn(img, float64(i*128), 0, 128, int8(i))
	}

	fi, err := os.Create("./testdata/" + file + ".png")
	a.NotError(err).NotNil(fi)
	defer func() {
		a.NotError(fi.Close())
	}()

	a.NotError(png.Encode(fi, img))
}

func TestBlocks(t *testing.T) {
	a := assert.New(t)

	for k, v := range blocks {
		drawImage("b"+strconv.Itoa(k), v, a)
	}
}
