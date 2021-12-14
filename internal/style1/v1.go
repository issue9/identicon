// SPDX-License-Identifier: MIT

// Package style1 风格 1 的头像
package style1

import (
	"image"
	"image/color"
)

// DrawBlocks 将九个方格都填上内容
//
// p 为画板；
// c 为中间方格的填充函数；
// b1、b2 为边上 8 格的填充函数；
// b1Angle 和 b2Angle 为 b1、b2 的起始旋转角度。
func DrawBlocks(p *image.Paletted, size, c, b1, b2, b1Angle, b2Angle int, fc color.Color) {
	b1 = b1 % len(blocks)
	b2 = b2 % len(blocks)
	c = c % len(centerBlocks)
	b1Angle = b1Angle % 4
	b2Angle = b2Angle % 4

	cc := centerBlocks[c]
	bb1 := blocks[b1]
	bb2 := blocks[b2]

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

	cc(p, blockSize+padding, blockSize+padding, blockSize, 0)

	bb1(p, 0+padding, 0+padding, blockSize, b1Angle)
	bb2(p, blockSize+padding, 0+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	bb1(p, twoBlockSize+padding, 0+padding, blockSize, b1Angle)
	bb2(p, twoBlockSize+padding, blockSize+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	bb1(p, twoBlockSize+padding, twoBlockSize+padding, blockSize, b1Angle)
	bb2(p, blockSize+padding, twoBlockSize+padding, blockSize, b2Angle)

	b1Angle = incr(b1Angle)
	b2Angle = incr(b2Angle)
	bb1(p, 0+padding, twoBlockSize+padding, blockSize, b1Angle)
	bb2(p, 0+padding, blockSize+padding, blockSize, b2Angle)
}
