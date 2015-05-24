// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"math"
)

const (
	minSize = 16
)

// Identicon 用于产生统一颜色和尺寸的头像。
type Identicon struct {
	colors []color.Color
	size   int
}

// 声明一个Identicon实例。
// back, fore表示头像的背景和前景色。
// size表示整个头像的大小。
func New(back, fore color.Color, size int) (*Identicon, error) {
	if size < minSize {
		return nil, fmt.Errorf("New:产生的图片尺寸(%v)不能小于%v", size, minSize)
	}

	return &Identicon{
		colors: []color.Color{back, fore},
		size:   size,
	}, nil
}

// 根据data数据产生一张唯一性的头像图片。
func (i *Identicon) Make(data []byte) image.Image {
	p := image.NewPaletted(image.Rect(0, 0, i.size, i.size), i.colors)
	return drawImage(p, i.size, data)
}

// 根据data数据产生一张唯一性的头像图片。
// back, fore头像的背景和前景色。
// size 头像的大小。
func Make(back, fore color.Color, size int, data []byte) (image.Image, error) {
	if size < minSize {
		return nil, fmt.Errorf("New:产生的图片尺寸(%v)不能小于%v", size, minSize)
	}

	p := image.NewPaletted(image.Rect(0, 0, size, size), []color.Color{back, fore})

	return drawImage(p, size, data), nil
}

// 将data转换成图像画在p上面。
// size为画板的长和宽。
func drawImage(p *image.Paletted, size int, data []byte) image.Image {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)

	// 第一个方块
	index := int(math.Abs(float64(sum[0]+sum[1]+sum[2]+sum[3]))) % len(blocks)
	b1 := blocks[index]

	// 第二个方块
	index = int(math.Abs(float64(sum[4]+sum[5]+sum[6]+sum[7]))) % len(blocks)
	b2 := blocks[index]

	// 中间方块
	index = int(math.Abs(float64(sum[8]+sum[9]+sum[10]+sum[11]))) % len(centerBlocks)
	c := centerBlocks[index]

	// 旋转角度
	angle := int(math.Abs(float64(sum[12]+sum[13]+sum[14]+sum[15]))) % 4

	drawBlocks(p, size, c, b1, b2, angle)
	return p
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
