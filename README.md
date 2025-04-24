go-slugify
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft/go-slugify)](https://goreportcard.com/report/github.com/khulnasoft/go-slugify)

Make Pretty Slugs.

Installation
------------

```
go get -u github.com/khulnasoft/go-slugify
```

Install CLI tool:

```
go get -u github.com/khulnasoft/go-slugify/slugify
$ slugify "北京kožušček,abc"
bei-jing-kozuscek-abc
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/khulnasoft/go-slugify


Usage
------

```go
package main

import (
	"fmt"
	"github.com/khulnasoft/go-slugify"
)

func main() {
	slugifier := (&slugify.Slugifier{}).ToLower(false).InvalidChar("-").WordSeparator("-")
	s := "北京kožušček,abc"
	fmt.Println(slugifier.Slugify(s))
	// Output: bei-jing-kozuscek-abc
}
```
