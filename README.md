# About

🔹 This tool searches the comment, including the "TODO" by using Go

## Send API Request
```
> go run server.go
```

## Usage
```go
package search

import (
	"github.com/gin-gonic/gin"
	"golang_practice/constants"
	"golang_practice/engine"
)

func Search(context *gin.Context) {
	pattern := context.Param("pattern")
	packageName := context.Query("p")
	context.JSON(constants.HTTP_OK, engine.Search(packageName, pattern))
}
```

you can make sure that the http response is success
[![https://gyazo.com/a9a438a1687264a098d5f9454060110f](https://i.gyazo.com/a9a438a1687264a098d5f9454060110f.png)](https://gyazo.com/a9a438a1687264a098d5f9454060110f)

When you open google chrome, and then input

***`http://localhost:9000/api/search/todo?p=fmt`*** on the url field like below.
[![https://gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7](https://i.gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7.png)](https://gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7)

you could receive json format data like below

## Result
```go

[
  {
    Filepath: "/Users/yogoken/.gvm/gos/go1.6.2/src/fmt/format.go",
    Comments: [
      {
        Row: 332,
        Comment: "TODO: Avoid buffer by pre-padding. "
      }
    ]
  },
  {
    Filepath: "/Users/yogoken/.gvm/gos/go1.6.2/src/fmt/scan.go",
    Comments: [
      {
        Row: 747,
        Comment: "TODO: accept N and Ni independently? "
      }
    ]
  }
]
```

## Other result
[![https://gyazo.com/f8b8c6429a9b9b9f74726b7db8b1cfb2](https://i.gyazo.com/f8b8c6429a9b9b9f74726b7db8b1cfb2.gif)](https://gyazo.com/f8b8c6429a9b9b9f74726b7db8b1cfb2)
