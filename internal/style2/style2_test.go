// SPDX-License-Identifier: MIT

package style2

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"testing"

	"github.com/issue9/assert/v3"
)

var (
	back = color.RGBA{R: 255, G: 0, B: 0, A: 100}
	fore = color.RGBA{R: 0, G: 255, B: 255, A: 100}
	size = 128
)

func TestDraw(t *testing.T) {
	a := assert.New(t, false)
	p := []color.Color{back, fore}

	for i := 0; i < 20; i++ {
		img := image.NewPaletted(image.Rect(0, 0, size, size), p)
		a.NotNil(img)

		fi, err := os.Create("./testdata/v2-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)

		Draw(img, size/Blocks, uint32(123222243)|(uint32(i)+11133))
		a.NotError(png.Encode(fi, img))

		a.NotError(fi.Close()) // 关闭文件
	}
}
