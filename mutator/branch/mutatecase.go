package branch

import (
	"go/ast"
	"go/types"

	"github.com/zimmski/go-mutesting/mutator"
)

func init() {
	mutator.Register("branch/case", MutatorCase)
}

// MutatorCase implements a mutator for case clauses.
func MutatorCase(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	_ = "STUB: not implemented"
	return nil
}
