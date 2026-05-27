package branch

import (
	"go/ast"
	"go/types"

	"github.com/zimmski/go-mutesting/mutator"
)

func init() {
	mutator.Register("branch/if", MutatorIf)
}

// MutatorIf implements a mutator for if and else if branches.
func MutatorIf(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	_ = "STUB: not implemented"
	return nil
}
