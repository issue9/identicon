// SPDX-License-Identifier: MIT

package identicon

import (
	"fmt"
	"hash"
	"hash/fnv"
	"image"
	"image/color"
	"image/color/palette"
	"math"
	"math/rand"
	"strconv"

	"github.com/issue9/identicon/v2/internal/style1"
	"github.com/issue9/identicon/v2/internal/style2"
)

type Style int8

const (
	Style1 Style = iota + 1 // 旧版本风格
	Style2                  // Style2 风格，性能略高于 Style1
)

// Identicon 用于产生统一尺寸的头像
//
// 可以根据用户提供的数据，经过一定的算法，自动产生相应的图案和颜色。
type Identicon struct {
	style      Style
	foreColors []color.Color
	backColor  color.Color
	size       int
	rect       image.Rectangle
	hash       hash.Hash32

	// style v2
	bitsPerPoint int
}

// S1 采用 style1 风格的头像
//
// 背景为透明，前景由 image/color/palette.WebSafe 指定；
func S1(size int) *Identicon {
	return New(Style1, size, color.Transparent, palette.WebSafe...)
}

// S2 采用 style2 风格的头像
//
// 背景为透明，前景由 image/color/palette.WebSafe 指定；
func S2(size int) *Identicon {
	return New(Style2, size, color.Transparent, palette.WebSafe...)
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
		if size < style1.MinSize {
			panic(fmt.Sprintf("参数 size 的值 %d 不能小于 %d", size, style1.MinSize))
		}
	case Style2:
		if size%style2.Blocks != 0 {
			panic(fmt.Sprintf("参数 size 的值 %d 必须为 %d 的倍数", size, style2.Blocks))
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
		bitsPerPoint: size / style2.Blocks,
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

	fc := int(sum&0xf0_f0_f0_f0) % len(i.foreColors)
	p := image.NewPaletted(i.rect, []color.Color{i.backColor, i.foreColors[fc]})

	switch i.style {
	case Style1:
		style1.DrawBlocks(p, i.size, sum)
		return p
	case Style2:
		style2.Draw(p, i.size, i.bitsPerPoint, sum)
		return p
	default:
		panic("无效的 style")
	}
}

// Make 根据 data 数据产生一张唯一性的头像图片
//
// size 头像的大小。
// back, fore 头像的背景和前景色。
func Make(style Style, size int, back, fore color.Color, data []byte) image.Image {
	return New(style, size, back, fore).Make(data)
}
