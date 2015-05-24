// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identicon

var (
	cos = []float64{1, 0, -1, 0}
	sin = []float64{0, 1, 0, -1}
)

// 将points中的所有点，以x,y为原点旋转angle个角度。
// angle取值只能是[0,1,2,3]，分别表示[0，90，180，270]
func rotate(points []float64, x, y float64, angle int) {
	if angle > 3 {
		panic("angle必须0,1,2,3三值之一")
	}

	for i := 0; i < len(points); i += 2 {
		px := points[i]
		py := points[i+1]
		points[i] = (px-x)*cos[angle] - (py-y)*sin[angle] + x
		points[i+1] = (px-x)*sin[angle] + (py-y)*cos[angle] + y
	}
}

// 判断某个点是否在多边形之内，不包含构成多边形的线和点
// x,y 需要判断的点坐标
// points 组成多边形的所顶点
func pointInPolygon(x float64, y float64, points []float64) bool {
	if len(points) < 6 { // 顶点数量少于3个，肯定无法合并
		return false
	}

	// 大致算法如下：
	// 把整个平面以给定的测试点为原点分两部分:
	// - y>0，包含(x>0 && y==0)
	// - y<0，包含(x<0 && y==0)
	// 依次扫描每一个点，当该点与前一个点处于不同部分时（即一个在y>0区，一个在y<0区），
	// 则判断从前一点到当前点是顺时针还是逆时针（以给定的测试点为原点），如果是顺时针r++，否则r--。
	// 结果为：2==abs(r)。

	r := 0
	points = append(points, points[0], points[1]) // 将起始点放入尾部，形成一个闭合区域

	x1, y1 := points[0], points[1]
	prev := (y1 > y) || ((x1 > x) && (y1 == y))
	for i := 2; i < len(points); i += 2 {
		x2, y2 := points[i], points[i+1]
		curr := (y2 > y) || ((x2 > x) && (y2 == y))

		if curr == prev {
			x1, y1 = x2, y2
			continue
		}

		mul := (x1-x)*(y2-y) - (x2-x)*(y1-y)
		switch {
		case mul > 0:
			r++
		case mul < 0:
			r--
		}
		x1, y1 = x2, y2
		prev = curr
	}

	return r == 2 || r == -2
}
