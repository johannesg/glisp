package reader

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Reader struct {
	input string
	line  int
	start int
	pos   int
	width int
}

type Form interface {
}

type FormEmpty struct {
}

type FormError struct {
	Val string
}

type FormSymbol struct {
	Name string
}

type FormNumber struct {
	Val int
}

const eof = -1

func New(input string) *Reader {
	return &Reader{input: input}
}

func (reader *Reader) Read() Form {
	return reader.read(reader.next())
}

func (reader *Reader) read(r rune) Form {
	for ; r != eof; r = reader.next() {
		switch {
		case unicode.IsSpace(r):
			reader.ignore()
		case strings.ContainsRune("+-", r), unicode.IsNumber(r):
			return reader.readNumberOrSymbol()
		case strings.ContainsRune("*!_'?", r), unicode.IsLetter(r):
			return reader.readSymbol()
		default:
			return FormError{Val: reader.input[reader.start:reader.pos]}
		}
	}

	return FormEmpty{}
}

func (reader *Reader) readNumberOrSymbol() Form {
	for {
		r := reader.next()
		switch {
		case unicode.IsNumber(r):
			return reader.readNumber()
		default:
			return reader.readSymbol()
		}
	}
}

func (reader *Reader) readSymbol() Form {
	for {
		r := reader.next()
		switch {
		case strings.ContainsRune("*!_'?+-", r), unicode.IsLetter(r), unicode.IsNumber(r):
		default:
			reader.backup()
			return FormSymbol{Name: reader.input[reader.start:reader.pos]}
		}
	}
}

func (reader *Reader) readNumber() Form {
	for {
		r := reader.next()
		switch {
		case unicode.IsNumber(r):
		default:
			reader.backup()
			s := reader.input[reader.start:reader.pos]
			i, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return FormError{Val: s}
			}
			return FormNumber{Val: int(i)}
		}
	}
}

func (l *Reader) peek() (r rune) {
	r = l.next()
	l.backup()
	return
}

func (l *Reader) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *Reader) backup() {
	l.pos -= l.width
}

func (l *Reader) ignore() {
	l.start = l.pos
}
