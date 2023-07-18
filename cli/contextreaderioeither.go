package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	C "github.com/urfave/cli/v2"
)

func generateContextReaderIOEitherEitherize(f, fg *os.File, i int) {
	// non generic version
	fmt.Fprintf(f, "\n// Eitherize%d converts a function with %d parameters returning a tuple into a function with %d parameters returning a [ReaderIOEither[R]]\n// The inverse function is [Uneitherize%d]\n", i, i, i, i)
	fmt.Fprintf(f, "func Eitherize%d[F ~func(context.Context", i)
	for j := 0; j < i; j++ {
		fmt.Fprintf(f, ", T%d", j)
	}
	fmt.Fprintf(f, ") (R, error)")
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
	fmt.Fprintf(f, ") ReaderIOEither[R] {\n")
	fmt.Fprintf(f, "  return G.Eitherize%d[ReaderIOEither[R]](f)\n", i)
	fmt.Fprintln(f, "}")

	// generic version
	fmt.Fprintf(fg, "\n// Eitherize%d converts a function with %d parameters returning a tuple into a function with %d parameters returning a [GRA]\n// The inverse function is [Uneitherize%d]\n", i, i, i, i)
	fmt.Fprintf(fg, "func Eitherize%d[GRA ~func(context.Context) GIOA, F ~func(context.Context", i)
	for j := 0; j < i; j++ {
		fmt.Fprintf(fg, ", T%d", j)
	}
	fmt.Fprintf(fg, ") (R, error), GIOA ~func() E.Either[error, R]")
	for j := 0; j < i; j++ {
		fmt.Fprintf(fg, ", T%d", j)
	}
	fmt.Fprintf(fg, ", R any](f F) func(")
	for j := 0; j < i; j++ {
		if j > 0 {
			fmt.Fprintf(fg, ", ")
		}
		fmt.Fprintf(fg, "T%d", j)
	}
	fmt.Fprintf(fg, ") GRA {\n")
	fmt.Fprintf(fg, "  return RE.Eitherize%d[GRA](f)\n", i)
	fmt.Fprintln(fg, "}")
}

func generateContextReaderIOEitherHelpers(filename string, count int) error {
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
	// construct subdirectory
	genFilename := filepath.Join("generic", filename)
	err = os.MkdirAll("generic", os.ModePerm)
	if err != nil {
		return err
	}
	fg, err := os.Create(filepath.Clean(genFilename))
	if err != nil {
		return err
	}
	defer fg.Close()

	// log
	log.Printf("Generating code in [%s] for package [%s] with [%d] repetitions ...", filename, pkg, count)

	writePackage(f, pkg)

	fmt.Fprintf(f, `
import (
	"context"

	G "github.com/IBM/fp-go/context/%s/generic"
)
`, pkg)

	writePackage(fg, "generic")

	fmt.Fprintf(fg, `
import (
	"context"

	E "github.com/IBM/fp-go/either"
	RE "github.com/IBM/fp-go/readerioeither/generic"
)
`)

	generateContextReaderIOEitherEitherize(f, fg, 0)

	for i := 1; i <= count; i++ {
		// eitherize
		generateContextReaderIOEitherEitherize(f, fg, i)
	}

	return nil
}

func ContextReaderIOEitherCommand() *C.Command {
	return &C.Command{
		Name:  "contextreaderioeither",
		Usage: "generate code for ContextReaderIOEither",
		Flags: []C.Flag{
			flagCount,
			flagFilename,
		},
		Action: func(ctx *C.Context) error {
			return generateContextReaderIOEitherHelpers(
				ctx.String(keyFilename),
				ctx.Int(keyCount),
			)
		},
	}
}
