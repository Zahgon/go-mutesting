package astutil

import (
	"go/ast"
	"go/types"
)

// IdentifiersInStatement returns all identifiers with their found in a statement.
func IdentifiersInStatement(pkg *types.Package, info *types.Info, stmt ast.Stmt) []ast.Expr {
	_ = "STUB: not implemented"
	return nil
}

type identifierWalker struct {
	identifiers []ast.Expr
	pkg         *types.Package
	info        *types.Info
}

func checkForSelectorExpr(node ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (w *identifierWalker) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

// Ignore the blank identifier

// Ignore keywords

// We are only interested in variables

// FIXME instead of manually creating a new node, clone it and trim the node from its comments and position https://github.com/zimmski/go-mutesting/issues/49

// Check if we need to instantiate the expression

// FIXME we need to clone the node and trim comments and position recursively https://github.com/zimmski/go-mutesting/issues/49

// FIXME we need to clone the node and trim comments and position recursively https://github.com/zimmski/go-mutesting/issues/49

// Functions returns all found functions.
func Functions(n ast.Node) []*ast.FuncDecl { _ = "STUB: not implemented"; return nil }

type functionWalker struct {
	functions []*ast.FuncDecl
}

func (w *functionWalker) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}
