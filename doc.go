// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 一个基于hash值生成随机图像的包。
// 一般用于为用户首次注册时生成一个美观的头像。
//
// 在identicon中，把图像分成以下九个部分:
//  -------------
//  | 1 | 2 | 3 |
//  -------------
//  | 4 | 5 | 6 |
//  -------------
//  | 7 | 8 | 9 |
//  -------------
// 其中1、3、9、7为不同角度(依次增加90度)的同一张图片，
// 2、6、8、4也是如此，这样可以保持图像是对称的，比较美观。
// 5则单独使用一张图片。
//
//  // 根据用户访问的IP，为其生成一张头像
//  img, _ := identicon.Make(color.NRGBA{},color.NRGBA{}, 128, []byte("192.168.1.1"))
//  fi, _ := os.Create("/tmp/u1.png")
//  png.Encode(fi, img)
//  fi.Close()
//
//  // 或者
//  ii, _ := identicon.New(color.NRGBA{}, color.NGRGA{}, 128)
//  img := ii.Make([]byte("192.168.1.1"))
//  img = ii.Make([]byte("192.168.1.2"))
package identicon

const Version = "0.1.3.150524"
