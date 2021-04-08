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

	"github.com/issue9/assert"
)

var (
	back  = color.RGBA{255, 0, 0, 100}
	fore  = color.RGBA{0, 255, 255, 100}
	fores = []color.Color{color.Black, color.RGBA{200, 2, 5, 100}, color.RGBA{2, 200, 5, 100}}
	size  = 128
)

// 依次画出各个网络的图像。
func TestBlocks(t *testing.T) {
	p := []color.Color{back, fore}

	a := assert.New(t)

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
	a := assert.New(t)

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
	a := assert.New(t)

	for i := 0; i < 20; i++ {
		img, err := Make(size, back, fore, []byte("make-"+strconv.Itoa(i)))
		a.NotError(err).NotNil(img)

		fi, err := os.Create("./testdata/make-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Make(t *testing.T) {
	a := assert.New(t)

	ii, err := New(size, back, fores...)
	a.NotError(err).NotNil(ii)

	for i := 0; i < 20; i++ {
		img := ii.Make([]byte("identicon-" + strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/identicon-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Rand(t *testing.T) {
	a := assert.New(t)

	ii, err := New(size, back, fores...)
	a.NotError(err).NotNil(ii)
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

func BenchmarkMake(b *testing.B) {
	a := assert.New(b)
	for i := 0; i < b.N; i++ {
		img, err := Make(size, back, fore, []byte("Make"))
		a.NotError(err).NotNil(img)
	}
}

func BenchmarkIdenticon_Make(b *testing.B) {
	a := assert.New(b)

	ii, err := New(size, back, fores...)
	a.NotError(err).NotNil(ii)

	for i := 0; i < b.N; i++ {
		img := ii.Make([]byte("Make"))
		a.NotNil(img)
	}
}

func BenchmarkIdenticon_Rand(b *testing.B) {
	a := assert.New(b)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ii, err := New(size, back, fores...)
	a.NotError(err).NotNil(ii)

	for i := 0; i < b.N; i++ {
		img := ii.Rand(r)
		a.NotNil(img)
	}
}
