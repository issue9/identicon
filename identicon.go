// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
)

const (
	minSize       = 16 // 图片的最小尺寸
	maxForeColors = 32 // 在New()函数中可以指定的最大颜色数量
)

// Identicon 用于产生统一尺寸的头像。
// 可以根据用户提供的数据，经过一定的算法，自动产生相应的图案和颜色。
type Identicon struct {
	foreColors []color.Color
	backColor  color.Color
	size       int
}

// 声明一个Identicon实例。
// size表示整个头像的大小。
// back表示前景色。
// fore表示所有可能的前景色，会为每个图像随机挑选一个作为其前景色。不要与背景色太相近。
func New(size int, back color.Color, fore ...color.Color) (*Identicon, error) {
	if len(fore) == 0 || len(fore) > maxForeColors {
		return nil, fmt.Errorf("前景色数量必须介于[1]~[%v]之间，当前为[%v]", maxForeColors, len(fore))
	}

	if size < minSize {
		return nil, fmt.Errorf("参数size的值(%v)不能小于%v", size, minSize)
	}

	return &Identicon{
		foreColors: fore,
		backColor:  back,
		size:       size,
	}, nil
}

// 根据data数据产生一张唯一性的头像图片。
func (i *Identicon) Make(data []byte) image.Image {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)

	// 第一个方块
	index := abs(sum[0]+sum[1]+sum[2]+sum[3]) % len(blocks)
	b1 := blocks[index]

	// 第二个方块
	index = abs(sum[4]+sum[5]+sum[6]+sum[7]) % len(blocks)
	b2 := blocks[index]

	// 中间方块
	index = abs(sum[8]+sum[9]+sum[10]+sum[11]) % len(centerBlocks)
	c := centerBlocks[index]

	// 旋转角度
	angle := abs(sum[12]+sum[13]+sum[14]) % 4

	// 根据最后一个字段，获取前景颜色
	index = abs(sum[15]) % len(i.foreColors)

	// 画布坐标从0开始，其长度应该是size-1
	p := image.NewPaletted(image.Rect(0, 0, i.size-1, i.size-1), []color.Color{i.backColor, i.foreColors[index]})
	drawBlocks(p, i.size, c, b1, b2, angle)
	return p
}

// 根据data数据产生一张唯一性的头像图片。
// size 头像的大小。
// back, fore头像的背景和前景色。
func Make(size int, back, fore color.Color, data []byte) (image.Image, error) {
	if size < minSize {
		return nil, fmt.Errorf("参数size的值(%v)不能小于%v", size, minSize)
	}

	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)

	// 第一个方块
	index := abs(sum[0]+sum[1]+sum[2]+sum[3]) % len(blocks)
	b1 := blocks[index]

	// 第二个方块
	index = abs(sum[4]+sum[5]+sum[6]+sum[7]) % len(blocks)
	b2 := blocks[index]

	// 中间方块
	index = abs(sum[8]+sum[9]+sum[10]+sum[11]) % len(centerBlocks)
	c := centerBlocks[index]

	// 旋转角度
	angle := abs(sum[12]+sum[13]+sum[14]+sum[15]) % 4

	// 画布坐标从0开始，其长度应该是size-1
	p := image.NewPaletted(image.Rect(0, 0, size-1, size-1), []color.Color{back, fore})
	drawBlocks(p, size, c, b1, b2, angle)
	return p, nil
}

// 将九个方格都填上内容。
// p为画板。
// c为中间方格的填充函数。
// b1,b2为边上8格的填充函数。
// angle为b1,b2的起始旋转角度。
func drawBlocks(p *image.Paletted, size int, c, b1, b2 blockFunc, angle int) {
	blockSize := float64(size / 3) // 每个格子的长宽
	twoBlockSize := 2 * blockSize

	incr := func() { // 增加angle的值，但不会大于3
		if angle > 2 {
			angle = 0
		} else {
			angle++
		}
	}

	c(p, blockSize, blockSize, blockSize, 0)

	b1(p, 0, 0, blockSize, angle)
	b2(p, blockSize, 0, blockSize, angle)

	incr()
	b1(p, twoBlockSize, 0, blockSize, angle)
	b2(p, twoBlockSize, blockSize, blockSize, angle)

	incr()
	b1(p, twoBlockSize, twoBlockSize, blockSize, angle)
	b2(p, blockSize, twoBlockSize, blockSize, angle)

	incr()
	b1(p, 0, twoBlockSize, blockSize, angle)
	b2(p, 0, blockSize, blockSize, angle)
}

func abs(x byte) int {
	if x < 0 {
		return int(-x)
	}
	return int(x)
}
