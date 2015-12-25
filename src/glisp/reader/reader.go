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
		return a.readIdentifier(t)
	case t.typ == tokenNumber:
		n, err := strconv.ParseInt(t.val, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse number: %v", n)
		}
		return Number{Val: int(n)}, nil
	case t.typ == tokenString:
		return Literal{Val: t.val}, nil
	case t.typ == tokenQuote:
		return a.readQForm()
	case t.typ == tokenEOF:
		return nil, nil
	default:
		return nil, fmt.Errorf("Invalid token: %v", t)
	}
}

func (a *reader) readIdentifier(t token) (Form, error) {
	switch t.val {
	case "true":
		return Boolean(true), nil
	case "false":
		return Boolean(false), nil
	default:
		return Symbol{Name: t.val}, nil
	}
}

func (a *reader) readList() (Form, error) {
	l := &List{}

	for t := range a.lexer.tokens {
		switch {
		case isDelim(t, ")"):
			return l, nil
		default:
			i, err := a.readForm(t)
			if err != nil {
				return nil, err
			}

			l.Items = append(l.Items, i)
		}
	}
	return nil, fmt.Errorf("List not closed")
}

func (r *reader) readQForm() (Form, error) {
	f, err := r.readForm(<-r.lexer.tokens)
	if err != nil {
		return nil, err
	}

	return &QForm{Form: f}, nil
}

func isDelim(t token, v string) bool {
	return t.typ == tokenDelim && v == t.val
}
