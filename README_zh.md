# RESTest

## 安装 ego:
```
go get -u github.com/go-ego/ego  
```

## 安装 [Build-tools](https://github.com/go-ego/re)
```
go get -u github.com/go-ego/re 
```
### re new 

新建 ego web 项目 my-webapp

```
$ re new my-webapp
```

### re run

运行 my-webapp:
```
$ cd my-webapp && re run
```

## 测试 restful 接口:
```Go

package main

import (
	"github.com/go-ego/ego"
)

const httpUrl string = "http://127.0.0.1:3000"

func main() {

  router := ego.Classic()

  router.Static("/js", "./views/js")
  router.Static("/src", "./views/src")
  router.GlobHTML("views/html/*")

  strUrl := httpUrl + "/test/list"
  paramMap := ego.Map{
    "lon":  "10.1010101", // http 参数
    "lat":  "20.202020",
    "type": "1",
  }
  router.TestHtml(strUrl, paramMap) // http url, http 参数, args (可选参数): 默认为 "data"

  router.Run(":3100")
}

```
生成访问路径: http://127.0.0.1:3000/t/list

## 测试带文件上传的 restful 接口:
```Go

package main

import (
	"github.com/go-ego/ego"
)

const httpUrl string = "http://127.0.0.1:3000"

func main() {

  router := ego.Classic()

  router.Static("/js", "./views/js")
  router.Static("/src", "./views/src")
  router.GlobHTML("views/html/*")

  strUrl := httpUrl + "/test/list"
  paramMap := ego.Map{
    "lon":  "10.1010101", // http 参数
    "lat":  "20.202020", 
    "type": "1",
  }

  router.TestFile(strUrl, paranMap, "./file.go", "upfile") // http url, http 参数, 文件路径, 上传文件参数

  router.Run(":3100")
}

```
生成访问路径: http://127.0.0.1:3000/t/list