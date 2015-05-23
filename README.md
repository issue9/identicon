identicon [![Build Status](https://travis-ci.org/issue9/identicon.svg?branch=master)](https://travis-ci.org/issue9/identicon)
======

为用户产生漂亮的随机头像。

![screenhost.1](https://raw.github.com/issue9/identicon/master/screenshot/1.png)
![screenhost.2](https://raw.github.com/issue9/identicon/master/screenshot/2.png)
![screenhost.3](https://raw.github.com/issue9/identicon/master/screenshot/3.png)
![screenhost.4](https://raw.github.com/issue9/identicon/master/screenshot/4.png)
![screenhost.5](https://raw.github.com/issue9/identicon/master/screenshot/5.png)
![screenhost.6](https://raw.github.com/issue9/identicon/master/screenshot/6.png)
![screenhost.7](https://raw.github.com/issue9/identicon/master/screenshot/7.png)

```go
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
```

### 安装

```shell
go get github.com/issue9/identicon
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/identicon)
[![GoDoc](https://godoc.org/github.com/issue9/identicon?status.svg)](https://godoc.org/github.com/issue9/identicon)


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
