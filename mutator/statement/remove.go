package statement

import (
	"go/ast"
	"go/types"

	"github.com/zimmski/go-mutesting/mutator"
)

func init() {
	mutator.Register("statement/remove", MutatorRemoveStatement)
}

func checkRemoveStatement(node ast.Stmt) bool { _ = "STUB: not implemented"; return false }

// MutatorRemoveStatement implements a mutator to remove statements.
func MutatorRemoveStatement(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	_ = "STUB: not implemented"
	return nil
}
