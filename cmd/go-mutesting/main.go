package main

import (
	"go/ast"
	"go/token"
	"go/types"
	"os"

	"github.com/zimmski/go-mutesting/mutator"
	_ "github.com/zimmski/go-mutesting/mutator/branch"
	_ "github.com/zimmski/go-mutesting/mutator/expression"
	_ "github.com/zimmski/go-mutesting/mutator/statement"
)

const (
	returnOk = iota
	returnHelp
	returnBashCompletion
	returnError
)

type options struct {
	General struct {
		Debug                bool `long:"debug" description:"Debug log output"`
		DoNotRemoveTmpFolder bool `long:"do-not-remove-tmp-folder" description:"Do not remove the tmp folder where all mutations are saved to"`
		Help                 bool `long:"help" description:"Show this help message"`
		Verbose              bool `long:"verbose" description:"Verbose log output"`
	} `group:"General options"`

	Files struct {
		Blacklist []string `long:"blacklist" description:"List of MD5 checksums of mutations which should be ignored. Each checksum must end with a new line character."`
		ListFiles bool     `long:"list-files" description:"List found files"`
		PrintAST  bool     `long:"print-ast" description:"Print the ASTs of all given files and exit"`
	} `group:"File options"`

	Mutator struct {
		DisableMutators []string `long:"disable" description:"Disable mutator by their name or using * as a suffix pattern"`
		ListMutators    bool     `long:"list-mutators" description:"List all available mutators"`
	} `group:"Mutator options"`

	Filter struct {
		Match string `long:"match" description:"Only functions are mutated that confirm to the arguments regex"`
	} `group:"Filter options"`

	Exec struct {
		Exec    string `long:"exec" description:"Execute this command for every mutation (by default the built-in exec command is used)"`
		NoExec  bool   `long:"no-exec" description:"Skip the built-in exec command and just generate the mutations"`
		Timeout uint   `long:"exec-timeout" description:"Sets a timeout for the command execution (in seconds)" default:"10"`
	} `group:"Exec options"`

	Test struct {
		Recursive bool `long:"test-recursive" description:"Defines if the executer should test recursively"`
	} `group:"Test options"`

	Remaining struct {
		Targets []string `description:"Packages, directories and files even with patterns (by default the current directory)"`
	} `positional-args:"true" required:"true"`
}

func checkArguments(args []string, opts *options) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}

func debug(opts *options, format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func verbose(opts *options, format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func exitError(format string, args ...interface{}) int { _ = "STUB: not implemented"; return 0 }

type mutatorItem struct {
	Name    string
	Mutator mutator.Mutator
}

type mutationStats struct {
	passed     int
	failed     int
	duplicated int
	skipped    int
}

func (ms *mutationStats) Score() float64 { _ = "STUB: not implemented"; return 0 }

func (ms *mutationStats) Total() int { _ = "STUB: not implemented"; return 0 }

func mainCmd(args []string) int { _ = "STUB: not implemented"; return 0 }

func mutate(opts *options, mutators []mutatorItem, mutationBlackList map[string]struct{}, mutationID int, pkg *types.Package, info *types.Info, file string, fset *token.FileSet, src ast.Node, node ast.Node, tmpFile string, execs []string, stats *mutationStats) int {
	_ = "STUB: not implemented"
	return 0
}

// Ignore original state

func mutateExec(opts *options, pkg *types.Package, file string, src ast.Node, mutationFile string, execs []string) (execExitCode int) {
	_ = "STUB: not implemented"
	return 0
}

// Tests passed -> FAIL

// Tests failed -> PASS

// Did not compile -> SKIP

// Unknown exit code -> SKIP

// TODO timeout here

func main() {
	os.Exit(mainCmd(os.Args[1:]))
}

func saveAST(mutationBlackList map[string]struct{}, file string, fset *token.FileSet, node ast.Node) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}
