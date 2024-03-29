// SPDX-FileCopyrightText: 2015-2024 caixw
//
// SPDX-License-Identifier: MIT

// Package identicon 一个基于 hash 值生成随机图像的包
//
// identicon 并没有统一的标准，一般用于在用户注册时，
// 取用户的邮箱或是访问 IP 等数据(也可以是其它任何数据)，
// 进行 hash 运算，之后根据 hash 数据，产生一张图像，
// 这样即可以为用户产生一张独特的头像，又不会泄漏用户的隐藏。
//
// 提供了两种风格的头像：Style1 和 Style2。
//
// style1
//
// 在 style1 中，把图像分成以下九个部分:
//
//	-------------
//	| 1 | 2 | 3 |
//	-------------
//	| 4 | 5 | 6 |
//	-------------
//	| 7 | 8 | 9 |
//	-------------
//
// 其中 1、3、9、7 为不同角度(依次增加 90 度)的同一张图片，
// 2、6、8、4 也是如此，这样可以保持图像是对称的，比较美观。
// 5 则单独使用一张图片。
//
// style2
//
// 将用户内容计算出 32 位的 hash 值，以 4 bit 为一行，
// 值为 1 表示有前景色，为 0 表示没有背景色，同时镜像到右边。
//
//	// 根据用户访问的 IP ，为其生成一张头像
//	img := identicon.Make(Style2, 128, color.NRGBA{},color.NRGBA{}, []byte("192.168.1.1"))
//	fi, _ := os.Create("/tmp/u1.png")
//	png.Encode(fi, img)
//	fi.Close()
//
//	// 或者
//	ii := identicon.New(Style1, 128, color.NRGBA{}, color.NRGBA{}, color.NRGBA{})
//	img := ii.Make([]byte("192.168.1.1"))
//	img = ii.Make([]byte("192.168.1.2"))
//
// NOTE: go test 会在当前目录的 testdata 文件夹下产生大量的随机图片。
// 要运行测试，必须保证该文件夹是存在的，且有相应的写入权限。
package identicon
