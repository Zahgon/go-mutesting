package mutator

import (
	"go/ast"
	"go/types"
)

// Mutator defines a mutator for mutation testing by returning a list of possible mutations for the given node.
type Mutator func(pkg *types.Package, info *types.Info, node ast.Node) []Mutation

var mutatorLookup = make(map[string]Mutator)

// New returns a new mutator instance given the registered name of the mutator.
// The error return argument is not nil, if the name does not exist in the registered mutator list.
func New(name string) (Mutator, error) { _ = "STUB: not implemented"; return *new(Mutator), nil }

// List returns a list of all registered mutator names.
func List() []string { _ = "STUB: not implemented"; return nil }

// Register registers a mutator instance function with the given name.
func Register(name string, mutator Mutator) { _ = "STUB: not implemented"; return }
