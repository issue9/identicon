// SPDX-License-Identifier: MIT

package identicon

import (
	"image/png"
	"os"
	"strconv"
	"testing"

	"github.com/issue9/assert/v2"
)

func TestIdenticon_Make_v2(t *testing.T) {
	a := assert.New(t, false)

	ii := S2(size)
	a.NotNil(ii)

	for i := 0; i < 20; i++ {
		img := ii.Make([]byte("identicon-" + strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/v2-identicon-make" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}
