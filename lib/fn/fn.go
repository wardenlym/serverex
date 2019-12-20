package fn

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// ProtectCall 调用一个函数 f
// f 中引发的panic会被recover为一个error,并返回
// 可以使用在服务线程中,防止由于未知的panic导致进程退出(= try catch)
func ProtectCall(f func()) (err error) {

	defer func() {
		info := recover()
		if info != nil {
			err = errors.New(fmt.Sprint(info))
		}
	}()

	f()

	return err
}

func Compile(src []byte) error {

	fset := token.NewFileSet()
	a, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}

	for _, d := range a.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}

	ast.Inspect(a, func(n ast.Node) bool {
		fmt.Printf("%[1]T %[1]v\n", n)
		return true
	})

	return nil
}
