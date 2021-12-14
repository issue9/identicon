// SPDX-License-Identifier: MIT

package style2

import (
	"image"
	"math/bits"
)

func Draw(p *image.Paletted, size, bitsPerPoint int, sum uint32) image.Image {
	lines := matrix(sum)

	var yBase, xBase int
	for y := 0; y < 8; y++ {
		line := lines[y]
		for yy := 0; yy < bitsPerPoint; yy++ {
			for x := 0; x < 8; x++ {
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
	for i := 0; i < 8; i++ {
		vv := uint8((v >> (i * 4) & 0x0f))
		vv <<= 4
		ret[i] = vv + bits.Reverse8(vv)
	}
	return ret
}
