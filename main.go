// Copyright 2017 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

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
	router.TestHtml(strUrl, paramMap) // http url, http parameter

	router.Run(":3100")
}
