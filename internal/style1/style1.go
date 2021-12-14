// SPDX-License-Identifier: MIT

// Package style1 风格 1 的头像
package style1

import (
	"image"
	"image/color"
)

// DrawBlocks 将九个方格都填上内容
//
// sum 由 hash 计算出的随机数；
func DrawBlocks(p *image.Paletted, size int, sum uint32, fc color.Color) {
	b1 := int(sum&0x00_00_00_ff) % len(blocks)
	b2 := int(sum&0x00_00_ff_00) % len(blocks)
	c := int(sum&0x00_ff_00_00) % len(centerBlocks)
	b1Angle := int(sum&0x0f_00_00_00) % 4
	b2Angle := int(sum&0xf0_00_00_00) % 4

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
