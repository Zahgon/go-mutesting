package expression

import (
	"go/ast"
	"go/types"

	"github.com/zimmski/go-mutesting/mutator"
)

func init() {
	mutator.Register("expression/remove", MutatorRemoveTerm)
}

// MutatorRemoveTerm implements a mutator to remove expression terms.
func MutatorRemoveTerm(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	_ = "STUB: not implemented"
	return nil
}
