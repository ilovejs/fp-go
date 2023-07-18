package cli

import (
	"fmt"
	"os"
	"time"
)

func writePackage(f *os.File, pkg string) {
	// print package
	fmt.Fprintf(f, "package %s\n\n", pkg)
	// some header
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f, "// This file was generated by robots at")
	fmt.Fprintf(f, "// %s\n", time.Now())
}
