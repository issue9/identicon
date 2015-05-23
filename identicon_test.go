// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/issue9/assert"
)

// func Make(fore, back color.NRGBA, size int, data []byte) (image.Image, error) {
func TestMake(t *testing.T) {
	a := assert.New(t)

	img, err := Make(color.NRGBA{12, 200, 12, 100}, color.NRGBA{0, 0, 0, 100}, 64, []byte("xnotepad.com"))
	a.NotError(err).NotNil(img)

	fi, err := os.Create("./testdata/md5-" + "caixw.png")
	a.NotError(err).NotNil(fi)
	defer func() {
		a.NotError(fi.Close())
	}()

	a.NotError(png.Encode(fi, img))
}
