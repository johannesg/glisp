package parse

import (
	"fmt"
	"unicode/utf8"
)

type stateFn func(*lexer) stateFn

type lexer struct {
	input string
	start int
	pos   int
	width int
	items chan item
}

const eof = -1

func lex(input string) (l *lexer) {
	l = &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()

	return
}

func (l *lexer) run() {
	state := lexProgram
	for state != nil {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(i itemType) {
	l.items <- item{typ: i, val: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{typ: itemError, val: fmt.Sprintf(format, args)}
	return nil
}
