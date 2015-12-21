package parse

import "fmt"

type astNode interface {
}

type astError struct {
	desc string
}

type astList struct {
	items []astNode
}

type astSymbol struct {
	name string
}

type astLiteral struct {
	val string
}

type ast struct {
	lexer *lexer
	nodes []astNode
}

type parseStateFn func(a *ast) parseStateFn

func parse(input string) (a *ast) {
	l := lex(input)

	a = &ast{lexer: l}

	a.run()

	return
}

func (a *ast) emit(n astNode) {
	a.nodes = append(a.nodes, n)
}

func (a *ast) run() {
	a.nodes = append(a.nodes, a.parseForm(<-a.lexer.tokens))
}

func (a *ast) parseForm(t token) astNode {
	switch {
	case isDelim(t, "("):
		return a.parseList()
	case t.typ == tokenIdentifier:
		return astSymbol{name: t.val}
	case t.typ == tokenNumber:
		return astLiteral{val: t.val}
	case t.typ == tokenString:
		return astLiteral{val: t.val}
	default:
		return a.errorf("Invalid token: %v", t)
	}
}

func (a *ast) parseList() astNode {
	l := astList{}

	for t := range a.lexer.tokens {
		switch {
		case isDelim(t, ")"):
			return l
		default:
			l.items = append(l.items, a.parseForm(t))
		}
	}
	return a.errorf("List not closed")
}

func isDelim(t token, v string) bool {
	return t.typ == tokenDelim && v == t.val
}

func (a *ast) errorf(format string, args ...interface{}) astError {
	return astError{desc: fmt.Sprintf(format, args)}
}
