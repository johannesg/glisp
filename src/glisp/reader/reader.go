package reader

import (
	"fmt"
	"strconv"
)

type Reader interface {
	Read() (Form, error)
}

type reader struct {
	lexer *lexer
	nodes []Form
}

func NewReader(input string) Reader {
	l := lex(input)

	return &reader{lexer: l}
}

func (a *reader) Read() (Form, error) {
	return a.readForm(<-a.lexer.tokens)
}

func (a *reader) readForm(t token) (Form, error) {
	switch {
	case isDelim(t, "("):
		return a.readList()
	case t.typ == tokenIdentifier:
		return Symbol{name: t.val}, nil
	case t.typ == tokenNumber:
		n, err := strconv.ParseInt(t.val, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse number: %v", n)
		}
		return Number{val: int(n)}, nil
	case t.typ == tokenString:
		return Literal{val: t.val}, nil
	case t.typ == tokenEOF:
		return nil, nil
	default:
		return nil, fmt.Errorf("Invalid token: %v", t)
	}
}

func (a *reader) readList() (Form, error) {
	l := List{}

	for t := range a.lexer.tokens {
		switch {
		case isDelim(t, ")"):
			return l, nil
		default:
			i, err := a.readForm(t)
			if err != nil {
				return nil, err
			}

			l.items = append(l.items, i)
		}
	}
	return nil, fmt.Errorf("List not closed")
}

func isDelim(t token, v string) bool {
	return t.typ == tokenDelim && v == t.val
}
