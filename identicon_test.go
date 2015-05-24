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

var (
	back = color.RGBA{255, 0, 0, 100}
	fore = color.RGBA{0, 255, 255, 100}
	size = 128
)

// 依次画出各个网络的图像。
func TestBlocks(t *testing.T) {
	p := []color.Color{back, fore}

	a := assert.New(t)

	for k, v := range blocks {
		img := image.NewPaletted(image.Rect(0, 0, size*4, size), p) // 横向4张图片大小

		for i := 0; i < 4; i++ {
			v(img, float64(i*size), 0, float64(size), i)
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
		drawBlocks(p, size, centerBlocks[c], blocks[b1], blocks[b2], 0)

		fi, err := os.Create("./testdata/draw-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, p))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestMake(t *testing.T) {
	a := assert.New(t)

	for i := 0; i < 20; i++ {
		img, err := Make(back, fore, size, []byte("make-"+strconv.Itoa(i)))
		a.NotError(err).NotNil(img)

		fi, err := os.Create("./testdata/make-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon(t *testing.T) {
	a := assert.New(t)

	i, err := New(back, fore, size)
	a.NotError(err).NotNil(i)
	img := i.Make([]byte("192.168.1.1"))
	a.NotNil(img)

	fi, err := os.Create("./testdata/identicon.png")
	a.NotError(err).NotNil(fi)
	a.NotError(png.Encode(fi, img))
	a.NotError(fi.Close()) // 关闭文件
}

// BenchmarkMake    3000    336798 ns/op
func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img, err := Make(back, fore, size, []byte("Make"))
		if err != nil || img == nil {
			b.Error("BenchmarkMake:Make时发生错误")
		}
	}
}

//BenchmarkIdenticon_Make	    5000	    337151 ns/op
func BenchmarkIdenticon_Make(b *testing.B) {
	ii, err := New(back, fore, size)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		img := ii.Make([]byte("Make"))
		if img == nil {
			b.Error("BenchmarkMake:Make时发生错误")
		}
	}
}
