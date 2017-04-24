package main

import (
	// "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go/ast" // abstract structure tree
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const (
	PATTERN                  = "TODO" // 左のTODOをFIXMEにして、go run src/main.go reflectにすると引っかかるようになる。
	DB_TYPE                  = "sqlite3"
	DB_NAME                  = "go_api.db"
	DB_CONNECT_ERROR_MESSAGE = "failed to connect database"
	HELLO_WORLD_MESSAGE      = "Hello world!"
)

type Todo struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Path    string `json:"path"`
	Row     int    `json:"row"`
	Comment string `json:"comment"`
}

func NewTodo(name string) Todo {
	return Todo{
		Path:    name,
		Row:     1,
		Comment: "todo comment",
	}
}

func CreateTodos(todoNames []string) []Todo {
	todos := make([]Todo, len(todoNames))

	for i, name := range todoNames {
		todos[i] = NewTodo(name)
	}
	return todos
}

func importPkg(path, dir string) *build.Package {

	p, err := build.Import(path, dir, build.ImportComment)
	// fmt.Println(p) => &{/usr/....doc.go...}
	if err != nil {
		panic(err)
	}

	if p.BinaryOnly {
		return nil
	}

	if p.IsCommand() {
		return nil
	}

	return p
}

func extractTODO(fname string) {

	// token.NewFileSet関数によって取得することができ、go/parserパッケージのParse系の関数によってファイル情報を追加して行く
	fset := token.NewFileSet()
	// fmt.Println(fset) => &{{{0 0} 0 0 0 0} 1 [] <nil>}
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	// fmt.Println(cmap)
	for n, cgs := range cmap {
		f := fset.File(n.Pos())
		for _, cg := range cgs {
			t := cg.Text()
			if strings.Contains(t, PATTERN) {
				fmt.Printf("%s:%v:\n%s\n", fname, f.Position(cg.Pos()).Line, t)
			}
		}
	}
}

func init() {
	db, err := gorm.Open(DB_TYPE, DB_NAME)
	if err != nil {
		panic(DB_CONNECT_ERROR_MESSAGE)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Todo{})
	todoNames := []string{
		"Todo1",
		"Todo2",
		"Todo3",
	}

	todos := CreateTodos(todoNames)
	count := 0
	db.Table("todos").Count(&count)
	if count == 0 {
		for _, todo := range todos {
			db.Create(&todo)
		}
	}
}

func main() {
	// カレントディレクトリの取得と変更
	dir, err := os.Getwd()
	// fmt.Println(dir) => /Users/yogoken/code/go-lobal
	// fmt.Println(err) => <nil>
	if err != nil {
		panic(err)
	}

	// fmt.Println(os.Args[0]) => /var/folders/....../command-line-arguments/...hello
	// fmt.Println(os.Args[1]) => fmt
	p := importPkg(os.Args[1], dir)
	// fmt.Println(p) => &{/usr/....ftm...}
	// fmt.Println(p.GoFiles) => [doc.go format.go print.go scan.go]
	for _, f := range p.GoFiles {
		extractTODO(filepath.Join(p.Dir, f))
		// fmt.Println(filepath.Join(p.Dir, f))
	}
	//
	// // ORM
	// db, err := sql.Open("mysql", "root:@/database")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()
	//
	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	//gorm.Open
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(DB_CONNECT_ERROR_MESSAGE)
	}
	defer db.Close()

	// todoList取得
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, HELLO_WORLD_MESSAGE)
	})
	router.GET("/todos", func(c *gin.Context) {
		var (
			todos   []Todo
			jsonMap map[string]interface{} = make(map[string]interface{})
		)
		db.Find(&todos)
		jsonMap["todos"] = todos
		c.JSON(200, jsonMap)
	})
	router.Run(":8000")
}
