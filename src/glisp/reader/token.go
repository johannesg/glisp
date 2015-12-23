package reader

import "fmt"

type tokenType int

type token struct {
	typ tokenType
	val string
}

func (i token) String() string {
	switch i.typ {
	case tokenError:
		return i.val
	case tokenEOF:
		return "EOF"
	}

	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

const (
	tokenError = iota
	tokenEOF
	tokenDelim
	tokenNumber
	tokenString
	tokenIdentifier
)
