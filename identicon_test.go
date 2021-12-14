// SPDX-License-Identifier: MIT

package identicon

import (
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/issue9/assert/v2"
)

var (
	back = color.RGBA{R: 255, G: 0, B: 0, A: 100}
	fore = color.RGBA{R: 0, G: 255, B: 255, A: 100}
	size = 128
)

func TestMake(t *testing.T) {
	a := assert.New(t, false)

	for i := 0; i < 20; i++ {
		img := Make(Style1, size, back, fore, []byte("make-"+strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/v1-make-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Make_v1(t *testing.T) {
	a := assert.New(t, false)

	ii := S1(size)
	a.NotNil(ii)

	for i := 0; i < 20; i++ {
		img := ii.Make([]byte("identicon-" + strconv.Itoa(i)))
		a.NotNil(img)

		fi, err := os.Create("./testdata/v1-identicon-make" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}

func TestIdenticon_Rand_v1(t *testing.T) {
	a := assert.New(t, false)

	ii := S1(size)
	a.NotNil(ii)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 20; i++ {
		img := ii.Rand(r)
		a.NotNil(img)

		fi, err := os.Create("./testdata/v1-identicon-rand-" + strconv.Itoa(i) + ".png")
		a.NotError(err).NotNil(fi)
		a.NotError(png.Encode(fi, img))
		a.NotError(fi.Close()) // 关闭文件
	}
}
