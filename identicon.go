// SPDX-License-Identifier: MIT

package identicon

import (
	"fmt"
	"hash"
	"hash/fnv"
	"image"
	"image/color"
	"math"
	"math/bits"
	"math/rand"
	"strconv"
)

type Style int8

const (
	V1 Style = iota + 1 // 旧版本风格
	V2                  // v2 风格，性能略高于 V1
)

const v1MinSize = 24

// Identicon 用于产生统一尺寸的头像
//
// 可以根据用户提供的数据，经过一定的算法，自动产生相应的图案和颜色。
type Identicon struct {
	style      Style
	foreColors []color.Color
	backColor  color.Color
	size       int
	rect       image.Rectangle

	// style v2
	hash         hash.Hash32
	bitsPerPoint int
}

// New 声明一个 Identicon 实例
//
// style 图片风格；
// size 头像的大小，应该将 size 的值保持在能被 3 整除的偶数，图片才会平整；
// back 前景色；
// fore 所有可能的前景色，会为每个图像随机挑选一个作为其前景色。
func New(style Style, size int, back color.Color, fore ...color.Color) *Identicon {
	if len(fore) == 0 {
		panic("必须指定 fore 参数")
	}

	switch style {
	case V1:
		if size < v1MinSize {
			panic(fmt.Sprintf("参数 size 的值 %d 不能小于 %d", size, v1MinSize))
		}
	case V2:
		if size%8 != 0 {
			panic(fmt.Sprintf("参数 size 的值 %d 必须为 8 的倍数", size))
		}
	}

	return &Identicon{
		style:      style,
		foreColors: fore,
		backColor:  back,
		size:       size,
		rect:       image.Rect(0, 0, size, size),

		// hash
		hash:         fnv.New32a(),
		bitsPerPoint: size / 8,
	}
}

// Rand 随机生成图案
func (i *Identicon) Rand(r *rand.Rand) image.Image {
	v := r.Int63n(math.MaxInt64)
	return i.Make([]byte(strconv.FormatInt(v, 10)))
}

// Make 根据 data 数据随机图片
func (i *Identicon) Make(data []byte) image.Image {
	i.hash.Write(data)
	sum := i.hash.Sum32()
	i.hash.Reset()

	switch i.style {
	case V1:
		return i.v1(sum)
	case V2:
		return i.v2(sum)
	default:
		panic("无效的 style")
	}
}

func (i *Identicon) v2(sum uint32) image.Image {
	lines := matrix(sum)
	fc := sum % uint32(len(i.foreColors))
	p := image.NewPaletted(i.rect, []color.Color{i.backColor, i.foreColors[fc]})

	var yBase, xBase int
	for y := 0; y < 8; y++ {
		line := lines[y]
		for yy := 0; yy < i.bitsPerPoint; yy++ {
			for x := 0; x < 8; x++ {
				var index uint8
				if value := uint8(0b1000_0000 >> x); value&line == value {
					index = 1
				}

				for xx := 0; xx < i.bitsPerPoint; xx++ {
					p.SetColorIndex(xBase+xx, yBase+yy, index)
				}

				xBase += i.bitsPerPoint
			}
			xBase = 0
		} // end yy
		yBase += i.bitsPerPoint
	}

	return p
}

func matrix(v uint32) []uint8 {
	ret := make([]uint8, 8)
	for i := 0; i < 8; i++ {
		vv := uint8((v >> (i * 4) & 0x0f))
		vv <<= 4
		ret[i] = vv + bits.Reverse8(vv)
	}
	return ret
}

func (i *Identicon) v1(sum uint32) image.Image {
	b1 := int(sum&0x00_00_00_ff) % len(blocks)
	b2 := int(sum&0x00_00_ff_00) % len(blocks)
	c := int(sum&0x00_ff_00_00) % len(centerBlocks)
	b1Angle := int(sum&0x0f_00_00_00) % 4
	b2Angle := int(sum&0xf0_00_00_00) % 4
	fc := int(sum&0xf0_f0_f0_f0) % len(i.foreColors)

	p := image.NewPaletted(i.rect, []color.Color{i.backColor, i.foreColors[fc]})
	drawBlocks(p, i.size, centerBlocks[c], blocks[b1], blocks[b2], b1Angle, b2Angle)
	return p
}

// Make 根据 data 数据产生一张唯一性的头像图片
//
// size 头像的大小。
// back, fore 头像的背景和前景色。
func Make(style Style, size int, back, fore color.Color, data []byte) image.Image {
	return New(style, size, back, fore).Make(data)
}

// 将九个方格都填上内容。
// p 为画板；
// c 为中间方格的填充函数；
// b1、b2 为边上 8 格的填充函数；
// b1Angle 和 b2Angle 为 b1、b2 的起始旋转角度。
func drawBlocks(p *image.Paletted, size int, c, b1, b2 blockFunc, b1Angle, b2Angle int) {
	incr := func(a int) int {
		if a >= 3 {
			return 0
		}
		a++
		return a
	}

	padding := (size % 6) / 2 // 不能除尽的，边上留白。

	blockSize := size / 3
	twoBlockSize := 2 * blockSize

	c(p, blockSize+padding, blockSize+padding, blockSize, 0)

	b1(p, 0+padding, 0+padding, blockSize, b1Angle)
	b2(p, blockSize+padding, 0+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	b1(p, twoBlockSize+padding, 0+padding, blockSize, b1Angle)
	b2(p, twoBlockSize+padding, blockSize+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	b1(p, twoBlockSize+padding, twoBlockSize+padding, blockSize, b1Angle)
	b2(p, blockSize+padding, twoBlockSize+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	b1(p, 0+padding, twoBlockSize+padding, blockSize, b1Angle)
	b2(p, 0+padding, blockSize+padding, blockSize, b2Angle)
}
