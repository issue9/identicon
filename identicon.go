// SPDX-License-Identifier: MIT

package identicon

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"math/rand"
)

const (
	minSize       = 16 // 图片的最小尺寸
	maxForeColors = 32 // 在New()函数中可以指定的最大颜色数量
)

// Identicon 用于产生统一尺寸的头像
//
// 可以根据用户提供的数据，经过一定的算法，自动产生相应的图案和颜色。
type Identicon struct {
	foreColors []color.Color
	backColor  color.Color
	size       int
	rect       image.Rectangle
}

// New 声明一个 Identicon 实例
//
// size 表示整个头像的大小；
// back 表示前景色；
// fore 表示所有可能的前景色，会为每个图像随机挑选一个作为其前景色。
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

		// 画布坐标从0开始，其长度应该是 size-1
		rect: image.Rect(0, 0, size, size),
	}, nil
}

// Make 根据 data 数据产生一张唯一性的头像图片
func (i *Identicon) Make(data []byte) image.Image {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)

	b1 := int(sum[0]+sum[1]+sum[2]+sum[3]) % len(blocks)
	b2 := int(sum[4]+sum[5]+sum[6]+sum[7]) % len(blocks)
	c := int(sum[8]+sum[9]+sum[10]+sum[11]) % len(centerBlocks)
	angle := int(sum[12]+sum[13]+sum[14]) % 4
	color := int(sum[15]) % len(i.foreColors)

	return i.render(c, b1, b2, angle, color)
}

// Rand 随机生成图案
func (i *Identicon) Rand(r *rand.Rand) image.Image {
	b1 := r.Intn(len(blocks))
	b2 := r.Intn(len(blocks))
	c := r.Intn(len(centerBlocks))
	angle := r.Intn(4)
	color := r.Intn(len(i.foreColors))

	return i.render(c, b1, b2, angle, color)
}

func (i *Identicon) render(c, b1, b2, angle, foreColor int) image.Image {
	p := image.NewPaletted(i.rect, []color.Color{i.backColor, i.foreColors[foreColor]})
	drawBlocks(p, i.size, centerBlocks[c], blocks[b1], blocks[b2], angle)
	return p
}

// Make 根据 data 数据产生一张唯一性的头像图片
//
// size 头像的大小。
// back, fore头像的背景和前景色。
func Make(size int, back, fore color.Color, data []byte) (image.Image, error) {
	i, err := New(size, back, fore)
	if err != nil {
		return nil, err
	}
	return i.Make(data), nil
}

// 将九个方格都填上内容。
// p 为画板；
// c 为中间方格的填充函数；
// b1、b2 为边上 8 格的填充函数；
// angle 为 b1、b2 的起始旋转角度。
func drawBlocks(p *image.Paletted, size int, c, b1, b2 blockFunc, angle int) {
	// 每个格子的长宽。先转换成 float，再计算！
	blockSize := float64(size) / 3
	twoBlockSize := 2 * blockSize

	c(p, blockSize, blockSize, blockSize, 0)

	b1(p, 0, 0, blockSize, 0)
	b2(p, blockSize, 0, blockSize, 0)

	b1(p, twoBlockSize, 0, blockSize, 1)
	b2(p, twoBlockSize, blockSize, blockSize, 1)

	b1(p, twoBlockSize, twoBlockSize, blockSize, 2)
	b2(p, blockSize, twoBlockSize, blockSize, 2)

	b1(p, 0, twoBlockSize, blockSize, 3)
	b2(p, 0, blockSize, blockSize, 3)
}
