package branch

import (
	"go/ast"
	"go/types"

	"github.com/zimmski/go-mutesting/mutator"
)

func init() {
	mutator.Register("branch/else", MutatorElse)
}

// MutatorElse implements a mutator for else branches.
func MutatorElse(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	_ = "STUB: not implemented"
	return nil
}

// We ignore else ifs and nil blocks
