package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
)

var lineLimit int
var stmtLimit int

func init() {
	flag.IntVar(&stmtLimit, `s`, 20, `The maximum number of statements allowed in a function`)
	flag.IntVar(&lineLimit, `l`, 35, `The maximum number of lines allowed in a function`)
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println(`Usage:`, os.Args[0], `[options] target`)
		os.Exit(2)
	}
}

var fset = token.NewFileSet()

func main() {
	pkgs, err := parser.ParseDir(fset, flag.Arg(0), nil, 0)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, f := range file.Decls {
				decl, ok := f.(*ast.FuncDecl)
				if !ok {
					continue
				}
				_ = makeStmtMessage(decl.Name, parseStmts(decl.Body.List)) ||
					makeLineMessage(decl.Name, getLines(decl))
			}
		}
	}
}

func makeLineMessage(funcInfo *ast.Ident, lines int) bool {
	if lines <= lineLimit {
		return false
	}
	fmt.Printf("%v:Function is too long (%d > %d)\n",
		fset.Position(funcInfo.Pos()),
		lines, lineLimit)
	return true
}

func makeStmtMessage(funcInfo *ast.Ident, stmts int) bool {
	if stmts <= stmtLimit {
		return false
	}
	fmt.Printf("%v:Function has too many statements (%d > %d)\n",
		fset.Position(funcInfo.Pos()),
		stmts, stmtLimit)
	return true
}

func getLines(f *ast.FuncDecl) int {
	return fset.Position(f.End()).Line - fset.Position(f.Pos()).Line - 2
}

func parseStmts(s []ast.Stmt) (total int) {
	for _, v := range s {
		total++
		switch stmt := v.(type) {
		case *ast.BlockStmt:
			total += parseStmts(stmt.List) - 1
		case *ast.ForStmt, *ast.RangeStmt, *ast.IfStmt,
			*ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt:
			total += parseBodyListStmts(stmt)
		case *ast.CaseClause:
			total += parseStmts(stmt.Body)
		case *ast.AssignStmt:
			total += checkInlineFunc(stmt.Rhs[0])
		case *ast.GoStmt:
			total += checkInlineFunc(stmt.Call.Fun)
		case *ast.DeferStmt:
			total += checkInlineFunc(stmt.Call.Fun)
		}
	}
	return
}

func checkInlineFunc(stmt ast.Expr) int {
	if block, ok := stmt.(*ast.FuncLit); ok {
		return parseStmts(block.Body.List)
	}
	return 0
}

func parseBodyListStmts(t interface{}) int {
	i := reflect.ValueOf(t).Elem().FieldByName(`Body`).Elem().FieldByName(`List`).Interface()
	return parseStmts(i.([]ast.Stmt))
}
