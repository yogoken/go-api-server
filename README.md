# About
This tool searches the comment, including the "TODO" from the specified package of Go language.

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

When you open google chrome, and then input `http://localhost:9000/api/search/todo?p=fmt` on the url field like below.
[![https://gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7](https://i.gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7.png)](https://gyazo.com/b2a94b13b54070e046f88e6b5e58e6c7)

you can make sure that the http response is success
[![https://gyazo.com/a9a438a1687264a098d5f9454060110f](https://i.gyazo.com/a9a438a1687264a098d5f9454060110f.png)](https://gyazo.com/a9a438a1687264a098d5f9454060110f)

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


--------------------------------------------------------------------------------------------




# Premise

In one development team, we wanted to create a list of TODO comment that was written in the source code, I made a command-line tool.

This tool searches the comment, including the "TODO" from the package of the specified Go, is intended to be displayed on the standard output.

[![https://gyazo.com/9ba6056b6fa83fbbc6ad987a0131f10f](https://i.gyazo.com/9ba6056b6fa83fbbc6ad987a0131f10f.png)](https://gyazo.com/9ba6056b6fa83fbbc6ad987a0131f10f)

For example, you can search for TODO comments fmt package as follows.

The source code for this tool, please refer to the main.go that have been attached.

It should be noted, main.go that is attached operates in Go1.7 more.

# Task
In this team, we wanted the above tools and extended as follows.

```
1. Make it possible to change the character string to search (such as "TODO" or "FIXME")
2. Make it work with Go 1.6.2 or late
3. Offer server mode
4. If you start as a server mode, it operates as an API server
5.
To this server,
if you send HTTP request with 1. import path (such as fmt and net/http) of a package and
                              2. the character string to search                      
you can get the results in JSON format

6. At least in the response, which contains 1. the corresponding file path, 2. row, 3. the target comment
```

No.5 meant

```
> go run main.go fmt "TODO"
> POST/json HTTP/1.1........
>
{
    “path”: “/path/to/package/file.go”,
    “output”: [
        {
            “row”: 100
            “comment”: “TODO: abc”
        },
        {
            “row”: 200
            “comment”: “TODO: abc”
        }
    ]
}
```

So as to satisfy the above requirements, please correct the main.go.

It should be noted that the format of the response and request of the API does not matter if I have freely design.

It does not need to fit in a single package or a single file.

There is no need to comply with the writing of the appended to have main.go, it does not matter if I have freely improved.

In addition, we ask you to please answer the following in mind:.

- Premise that this tool has been maintained in the team
- It is actively improvement that can be improved in the existing code
- Attach a code that operation can be verified
- API documentation can be generated from the source code who created
