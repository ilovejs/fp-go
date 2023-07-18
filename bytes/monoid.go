package bytes

import (
	"bytes"

	A "github.com/ibm/fp-go/array"
	O "github.com/ibm/fp-go/ord"
)

var (
	// monoid for byte arrays
	Monoid = A.Monoid[byte]()

	// ConcatAll concatenates all bytes
	ConcatAll = A.ArrayConcatAll[byte]

	// Ord implements the default ordering on bytes
	Ord = O.MakeOrd(bytes.Compare, bytes.Equal)
)