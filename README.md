identicon [![Build Status](https://travis-ci.org/issue9/identicon.svg?branch=master)](https://travis-ci.org/issue9/identicon)
======

```go
type User struct {
    // 对应表中的id字段，为自增列，从0开始
    Id          int64      `identicon:"name(id);ai(0);"`
    // 对应表中的first_name字段，为索引index_name的一部分
    FirstName   string     `identicon:"name(first_name);index(index_name)"`
    LastName    string     `identicon:"name(first_name);index(index_name)"`
}

// 创建User表
e.Create(&User{})

// 更新id为1的记录
e.Update(&User{Id:1,FirstName:"abc"})
e.Where("id=?", 1).Table("#tbl_name").Update(true, "FirstName", "abc")

// 删除id为1的记录
e.Delete(&User{Id:1})
e.Where("id=?", 1).Table("#tbl_name").Delete(true, []interface{}{"id":1})

// 插入数据
e.Insert(&User{FirstName:"abc"})

// 查找数据
maps,err := e.Where("id<?", 5).Table("#tbl_name").SelectMap(true, "*")
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
