# RESTest

[简体中文](https://github.com/go-ego/RESTest/blob/master/README_zh.md)

## Install ego:
```
go get -u github.com/go-ego/ego  
```

## Install [Build-tools](https://github.com/go-ego/re)
```
go get -u github.com/go-ego/re 
```
### re new 
To create a new Ego web application

```
$ re new my-webapp
```

### re run

To run the application we just created, you can navigate to the application folder and execute:
```
$ cd my-webapp && re run
```

## Test restful api:
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
    "lon":  "10.1010101", // http parameter
    "lat":  "20.202020",
    "type": "1",
  }
  router.TestHtml(strUrl, paramMap) // http url, http parameter, args (optional parameters): The default is "data".

  router.Run(":3100")
}

```

Generate access path: http://127.0.0.1:3000/t/list

## Test the file upload with the restful api:
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
    "lon":  "10.1010101", // http parameter
    "lat":  "20.202020",
    "type": "1",
  }

  router.TestFile(fileUrl, paranMap, "./file.go", "upfile") // http url, http parameter, file path, upload file parameters

  router.Run(":3100")
}

```

Generate access path: http://127.0.0.1:3000/t/list