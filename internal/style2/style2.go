// SPDX-License-Identifier: MIT

// Package style2 风格 2 的头像
package style2

import (
	"image"
	"math/bits"
)

const Blocks = 8

const half = Blocks / 2

func Draw(p *image.Paletted, bitsPerPoint int, sum uint32) image.Image {
	lines := matrix(sum)

	var yBase, xBase int
	for y := 0; y < Blocks; y++ {
		line := lines[y]
		for yy := 0; yy < bitsPerPoint; yy++ {
			for x := 0; x < Blocks; x++ {
				var index uint8
				if value := uint8(0b1000_0000 >> x); value&line == value {
					index = 1
				}

				for xx := 0; xx < bitsPerPoint; xx++ {
					p.SetColorIndex(xBase+xx, yBase+yy, index)
				}

				xBase += bitsPerPoint
			}
			xBase = 0
		} // end yy
		yBase += bitsPerPoint
	}

	return p
}

func matrix(v uint32) []uint8 {
	ret := make([]uint8, 8)
	var size int
	for i := 0; i < Blocks; i++ {
		vv := uint8((v >> size & 0x0f)) << half
		ret[i] = vv + bits.Reverse8(vv)
		size += half
	}
	return ret
}
