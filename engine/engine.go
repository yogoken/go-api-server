package engine

import (
	"os"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"path/filepath"
	"golang_practice/util"
)

type CommentStruct struct {
	Row int
	Comment string
}

type ResultStruct struct {
	Filepath string
	Comments []CommentStruct
}

func Search(packageName string, pattern string) []ResultStruct {

	condition := strings.ToUpper(pattern)
	
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	results := make([]ResultStruct, 0)
	p := util.ImportPkg(packageName, dir)

	for _, file := range p.GoFiles {
		fname := filepath.Join(p.Dir, file)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		comments := make([]CommentStruct, 0)
		cmap := ast.NewCommentMap(fset, f, f.Comments)

		for n, cgs := range cmap {
			f := fset.File(n.Pos())
			for _, cg := range cgs {
				text := cg.Text()
				if strings.Contains(text, condition) {
					c := CommentStruct{Row: f.Position(cg.Pos()).Line, Comment: text}
					comments = append(comments, c)
				}
			}
		}

		if len(comments) > 0 {
			r := ResultStruct{Filepath: fname, Comments: comments}
			results = append(results, r)
		}
	}

	return results
}