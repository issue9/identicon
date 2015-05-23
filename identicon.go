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
	maxSize = 256
	minSize = 16
)

type Identicon struct {
	fore, back     color.Color
	size, gridSize int
}

// size九宫格中每个格子的像素
func New(foreground, background color.Color, gridSize int) (*Identicon, error) {
	size := gridSize * 3
	if size > maxSize || size < minSize {
		return nil, fmt.Errorf("New:产生的图片尺寸(%v)不能大于%v也不能小于%v", size, maxSize, minSize)
	}

	return &Identicon{
		fore:     foreground,
		back:     background,
		size:     size,
		gridSize: gridSize,
	}, nil
}

func (i *Identicon) Make(data []byte) (image.Image, error) {
	return nil, nil
}

func Make(fore, back color.NRGBA, size int, data []byte) (image.Image, error) {
	h := md5.New()
	h.Write(data)
	sum := h.Sum(nil)
	g1 := blocks[int(sum[0]+sum[1]+sum[2]+sum[3])%len(blocks)]                 // 第一个方块
	g2 := blocks[int(sum[4]+sum[5]+sum[6]+sum[7])%len(blocks)]                 // 第二个方块
	c := centerBlocks[int(sum[8]+sum[9]+sum[10]+sum[11])%len(centerBlocks)]    // 中间方块
	angle := int8(math.Abs(float64(int(sum[12]+sum[13]+sum[14]+sum[15]) % 4))) // 旋转角度

	p := image.NewPaletted(image.Rect(0, 0, size, size), []color.Color{back, fore})
	gridSize := float64(size / 3)
	c(p, gridSize, gridSize, gridSize, 0)

	angle = angleIncr(angle)
	g1(p, 0, 0, gridSize, angle)
	g2(p, gridSize, 0, gridSize, angle)

	angle = angleIncr(angle)
	g1(p, 2*gridSize, 0, gridSize, angleIncr(angle))
	g2(p, 2*gridSize, gridSize, gridSize, angleIncr(angle))

	angle = angleIncr(angle)
	g1(p, 2*gridSize, 2*gridSize, gridSize, angleIncr(angle))
	g2(p, gridSize, 2*gridSize, gridSize, angleIncr(angle))

	angle = angleIncr(angle)
	g1(p, 0, 2*gridSize, gridSize, angleIncr(angle))
	g2(p, 0, gridSize, gridSize, angleIncr(angle))

	return p, nil
}

func angleIncr(angle int8) int8 {
	if angle > 2 { // 如果已经为3，必须重置为0
		return 0
	}
	return angle + 1
}
