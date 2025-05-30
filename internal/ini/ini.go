package ini

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"io"
	"os"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
)

// OpenFile takes a path to a given file, and will open  and parse
// that file.
func OpenFile(path string) (Sections, error) {
	f, err := os.Open(path)
	if err != nil {
		return Sections{}, bytepluserr.New(ErrCodeUnableToReadFile, "unable to open file", err)
	}
	defer f.Close()

	return Parse(f)
}

// Parse will parse the given file using the shared config
// visitor.
func Parse(f io.Reader) (Sections, error) {
	tree, err := ParseAST(f)
	if err != nil {
		return Sections{}, err
	}

	v := NewDefaultVisitor()
	if err = Walk(tree, v); err != nil {
		return Sections{}, err
	}

	return v.Sections, nil
}

// ParseBytes will parse the given bytes and return the parsed sections.
func ParseBytes(b []byte) (Sections, error) {
	tree, err := ParseASTBytes(b)
	if err != nil {
		return Sections{}, err
	}

	v := NewDefaultVisitor()
	if err = Walk(tree, v); err != nil {
		return Sections{}, err
	}

	return v.Sections, nil
}
