// SPDX-License-Identifier: MIT

package identicon

import (
	"math/rand"
	"testing"
	"time"

	"github.com/issue9/assert/v2"
)

func BenchmarkMake(b *testing.B) {
	a := assert.New(b, false)
	for i := 0; i < b.N; i++ {
		img := Make(Style1, size, back, fore, []byte("Make"))
		a.NotNil(img)
	}
}

func BenchmarkIdenticon_Make_v1(b *testing.B) {
	a := assert.New(b, false)

	ii := New(Style1, size, back, fores...)
	a.NotNil(ii)

	for i := 0; i < b.N; i++ {
		img := ii.Make([]byte("Make"))
		a.NotNil(img)
	}
}

func BenchmarkIdenticon_Rand_v2(b *testing.B) {
	a := assert.New(b, false)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ii := New(Style1, size, back, fores...)
	a.NotNil(ii)

	for i := 0; i < b.N; i++ {
		img := ii.Rand(r)
		a.NotNil(img)
	}
}

func BenchmarkIdenticon_Make_v2(b *testing.B) {
	a := assert.New(b, false)

	ii := New(Style2, size, back, fores...)
	a.NotNil(ii)

	for i := 0; i < b.N; i++ {
		img := ii.Make([]byte("Make"))
		a.NotNil(img)
	}
}
