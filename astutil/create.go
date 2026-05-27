package astutil

import (
	"go/ast"
	"go/types"
)

// CreateNoopOfStatement creates a syntactically safe noop statement out of a given statement.
func CreateNoopOfStatement(pkg *types.Package, info *types.Info, stmt ast.Stmt) ast.Stmt {
	_ = "STUB: not implemented"
	return *new(ast.Stmt)
}

// CreateNoopOfStatements creates a syntactically safe noop statement out of a given statement.
func CreateNoopOfStatements(pkg *types.Package, info *types.Info, stmts []ast.Stmt) ast.Stmt {
	_ = "STUB: not implemented"
	return *new(ast.Stmt)
}
