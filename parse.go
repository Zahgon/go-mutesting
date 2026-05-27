package mutesting

import (
	"go/ast"
	"go/token"
	"go/types"
)

// ParseFile parses the content of the given file and returns the corresponding ast.File node and its file set for positional information.
// If a fatal error is encountered the error return argument is not nil.
func ParseFile(file string) (*ast.File, *token.FileSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ParseSource parses the given source and returns the corresponding ast.File node and its file set for positional information.
// If a fatal error is encountered the error return argument is not nil.
func ParseSource(data interface{}) (*ast.File, *token.FileSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ParseAndTypeCheckFile parses and type-checks the given file, and returns everything interesting about the file.
// If a fatal error is encountered the error return argument is not nil.
func ParseAndTypeCheckFile(file string, flags ...string) (*ast.File, *token.FileSet, *types.Package, *types.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}
