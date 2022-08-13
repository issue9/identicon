// SPDX-License-Identifier: MIT

package style1

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

// 依次画出各个网络的图像。
func TestBlocks(t *testing.T) {
	a := assert.New(t, false)
	p := []color.Color{back, fore}

	for k, v := range blocks {
		img := image.NewPaletted(image.Rect(0, 0, size*4, size), p) // 横向 4 张图片大小

		for i := 0; i < 4; i++ {
			v(img, i*size, 0, size, i)
		}

		fi, err := os.Create("./testdata/block-" + strconv.Itoa(k) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

// 产生一组测试图片
func TestDrawBlocks(t *testing.T) {
	a := assert.New(t, false)

	for i := 0; i < 20; i++ {
		p := image.NewPaletted(image.Rect(0, 0, size, size), []color.Color{back, fore})
		DrawBlocks(p, size, uint32(11132323+i))

		fi, err := os.Create("./testdata/draw-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, p))
		a.NotError(fi.Close()) // 关闭文件
	}
}
