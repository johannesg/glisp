package main

import (
	"fmt"
	"unicode/utf8"
)

type stateFn func(*lexer) stateFn

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens chan token
}

const eof = -1

func lex(input string) (l *lexer) {
	l = &lexer{
		input:  input,
		tokens: make(chan token),
	}

	go l.run()

	return
}

func (l *lexer) run() {
	state := lexProgram
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(i tokenType) {
	l.tokens <- token{typ: i, val: l.input[l.start:l.pos]}
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
	l.tokens <- token{typ: tokenError, val: fmt.Sprintf(format, args)}
	return nil
}
