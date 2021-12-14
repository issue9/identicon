// SPDX-License-Identifier: MIT

package identicon

import (
	"fmt"
	"hash"
	"hash/fnv"
	"image"
	"image/color"
	"math"
	"math/rand"
	"strconv"

	"github.com/issue9/identicon/internal/style1"
)

type Style int8

const (
	Style1 Style = iota + 1 // 旧版本风格
	Style2                  // v2 风格，性能略高于 V1
)

const style1MinSize = 24

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

// Make 根据 data 数据产生一张唯一性的头像图片
//
// size 头像的大小。
// back, fore 头像的背景和前景色。
func Make(style Style, size int, back, fore color.Color, data []byte) image.Image {
	return New(style, size, back, fore).Make(data)
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
	case Style1:
		if size < style1MinSize {
			panic(fmt.Sprintf("参数 size 的值 %d 不能小于 %d", size, style1MinSize))
		}
	case Style2:
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
	case Style1:
		return i.style1(sum)
	case Style2:
		return i.style2(sum)
	default:
		panic("无效的 style")
	}
}

func (i *Identicon) style1(sum uint32) image.Image {
	fc := int(sum&0xf0_f0_f0_f0) % len(i.foreColors)
	p := image.NewPaletted(i.rect, []color.Color{i.backColor, i.foreColors[fc]})
	style1.DrawBlocks(p, i.size, sum, i.foreColors[fc])
	return p
}
