package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	C "github.com/urfave/cli/v2"
)

func optionHKT(typeA string) string {
	return fmt.Sprintf("Option[%s]", typeA)
}

func generateOptionTraverseTuple(f *os.File, i int) {
	generateTraverseTuple1(optionHKT, "")(f, i)
}

func generateOptionSequenceTuple(f *os.File, i int) {
	generateSequenceTuple1(optionHKT, "")(f, i)
}

func generateOptionSequenceT(f *os.File, i int) {
	generateSequenceT1(optionHKT, "")(f, i)
}

func generateOptionize(f *os.File, i int) {
	// Create the optionize version
	fmt.Fprintf(f, "\n// Optionize%d converts a function with %d parameters returning a tuple of a return value R and a boolean into a function with %d parameters returning an Option[R]\n", i, i, i)
	fmt.Fprintf(f, "func Optionize%d[F ~func(", i)
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	fmt.Fprintf(f, ") (R, bool)")
	for j := 0; j < i; j++ {
		fmt.Fprintf(f, ", T%d", j)
	}
	fmt.Fprintf(f, ", R any](f F) func(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	fmt.Fprintf(f, ") Option[R] {\n")
	fmt.Fprintf(f, "  return func(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "t%d T%d", j, j)
	}
	fmt.Fprintf(f, ") Option[R] {\n")
	fmt.Fprintf(f, "    return optionize(func() (R, bool) {\n")
	fmt.Fprintf(f, "      return f(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "t%d", j)
	}
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f, "    })")
	fmt.Fprintln(f, "  }")
	fmt.Fprintln(f, "}")
}

func generateUnoptionize(f *os.File, i int) {
	// Create the optionize version
	fmt.Fprintf(f, "\n// Unoptionize%d converts a function with %d parameters returning a tuple of a return value R and a boolean into a function with %d parameters returning an Option[R]\n", i, i, i)
	fmt.Fprintf(f, "func Unoptionize%d[F ~func(", i)
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	fmt.Fprintf(f, ") Option[R]")
	for j := 0; j < i; j++ {
		fmt.Fprintf(f, ", T%d", j)
	}
	fmt.Fprintf(f, ", R any](f F) func(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "T%d", j)
	}
	fmt.Fprintf(f, ") (R, bool) {\n")
	fmt.Fprintf(f, "  return func(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "t%d T%d", j, j)
	}
	fmt.Fprintf(f, ") (R, bool) {\n")
	fmt.Fprintf(f, "    return Unwrap(f(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "t%d", j)
	}
	fmt.Fprintln(f, "))")
	fmt.Fprintln(f, "  }")
	fmt.Fprintln(f, "}")
}

func generateOptionHelpers(filename string, count int) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	pkg := filepath.Base(absDir)
	f, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	defer f.Close()
	// log
	log.Printf("Generating code in [%s] for package [%s] with [%d] repetitions ...", filename, pkg, count)

	// some header
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f, "// This file was generated by robots at")
	fmt.Fprintf(f, "// %s\n", time.Now())

	fmt.Fprintf(f, "package %s\n\n", pkg)

	fmt.Fprintf(f, `
import (
	A "github.com/ibm/fp-go/internal/apply"
	T "github.com/ibm/fp-go/tuple"
)
`)

	// print out some helpers
	fmt.Fprintf(f, `// optionize converts a nullary function to an option
func optionize[R any](f func() (R, bool)) Option[R] {
	if r, ok := f(); ok {
		return Some(r)
	}
	return None[R]()
}
`)

	// zero level functions

	// optionize
	generateOptionize(f, 0)
	// unoptionize
	generateUnoptionize(f, 0)

	for i := 1; i <= count; i++ {
		// optionize
		generateOptionize(f, i)
		// unoptionize
		generateUnoptionize(f, i)
		// sequenceT
		generateOptionSequenceT(f, i)
		// sequenceTuple
		generateOptionSequenceTuple(f, i)
		// traverseTuple
		generateOptionTraverseTuple(f, i)
	}

	return nil
}

func OptionCommand() *C.Command {
	return &C.Command{
		Name:  "option",
		Usage: "generate code for Option",
		Flags: []C.Flag{
			flagCount,
			flagFilename,
		},
		Action: func(ctx *C.Context) error {
			return generateOptionHelpers(
				ctx.String(keyFilename),
				ctx.Int(keyCount),
			)
		},
	}
}
