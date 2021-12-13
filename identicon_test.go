// SPDX-License-Identifier: MIT

package identicon

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/issue9/assert/v2"
)

var (
	back  = color.RGBA{R: 255, G: 0, B: 0, A: 100}
	fore  = color.RGBA{R: 0, G: 255, B: 255, A: 100}
	fores = []color.Color{color.Black, color.RGBA{R: 200, G: 2, B: 5, A: 100}, color.RGBA{R: 2, G: 200, B: 5, A: 100}}
	size  = 128
)

// 依次画出各个网络的图像。
func TestBlocks(t *testing.T) {
	p := []color.Color{back, fore}

	a := assert.New(t, false)

	for k, v := range blocks {
		img := image.NewPaletted(image.Rect(0, 0, size*4, size), p) // 横向4张图片大小

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
		c := (i + 1) % len(centerBlocks)
		b1 := (i + 2) % len(blocks)
		b2 := (i + 3) % len(blocks)
		drawBlocks(p, size, centerBlocks[c], blocks[b1], blocks[b2], 0, 0)

		fi, err := os.Create("./testdata/draw-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, p))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestMake(t *testing.T) {
	a := assert.New(t, false)

	for i := 0; i < 20; i++ {
		img := Make(V1, size, back, fore, []byte("make-"+strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/make-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Make_v1(t *testing.T) {
	a := assert.New(t, false)

	ii := New(V1, size, back, fores...)
	a.NotNil(ii)

	for i := 0; i < 20; i++ {
		img := ii.Make([]byte("identicon-" + strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/identicon-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Rand_v1(t *testing.T) {
	a := assert.New(t, false)

	ii := New(V1, size, back, fores...)
	a.NotNil(ii)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 20; i++ {
		img := ii.Rand(r)
		a.NotNil(img)

		fi, err := os.Create("./testdata/rand-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Make_v2(t *testing.T) {
	a := assert.New(t, false)

	ii := New(V2, size, back, fores...)
	a.NotNil(ii)

	for i := 20; i < 50; i++ {
		img := ii.Make([]byte("identicon-" + strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/v2-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}
