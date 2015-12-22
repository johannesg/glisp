package reader

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Reader struct {
	input   string
	line    int
	start   int
	pos     int
	width   int
	current rune
}

type Form interface {
}

type FormEmpty struct {
}

type FormError struct {
	Val  string
	Desc string
}

type FormSymbol struct {
	Name string
}

type FormNumber struct {
	Val int
}

type FormList struct {
	Items []Form
}

const eof = -1

func New(input string) *Reader {
	return &Reader{input: input}
}

func (reader *Reader) Read() Form {
	reader.next()

	return reader.read()
}

func (reader *Reader) read() Form {
	for reader.current != eof && unicode.IsSpace(reader.current) {
		reader.ignore()
		reader.next()
	}
	r := reader.current
	// log.Printf("Form: Current rune: %c", reader.current)

	switch {
	case r == eof:
		return FormEmpty{}
	case strings.ContainsRune("+-", r):
		return reader.readNumberOrSymbol()
	case strings.ContainsRune("*!_'?", r), unicode.IsLetter(r):
		return reader.readSymbol()
	case unicode.IsNumber(r):
		return reader.readNumber()
	case r == '(':
		return reader.readList()
	case strings.ContainsRune(")]}", r):
		return nil
	}
	return reader.errorf("Invalid token")
}

func (reader *Reader) readNumberOrSymbol() Form {
	for {
		reader.next()
		r := reader.current
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
		reader.next()
		r := reader.current
		switch {
		case strings.ContainsRune("*!_'?+-", r), unicode.IsLetter(r), unicode.IsNumber(r):
		default:
			return FormSymbol{Name: reader.input[reader.start : reader.pos-reader.width]}
		}
	}
}

func (reader *Reader) readNumber() Form {
	for {
		reader.next()
		r := reader.current
		switch {
		case unicode.IsNumber(r):
		default:
			s := reader.input[reader.start : reader.pos-reader.width]
			i, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return FormError{Val: s}
			}
			return FormNumber{Val: int(i)}
		}
	}
}

func (reader *Reader) readList() Form {
	reader.ignore()
	reader.next()
	// log.Printf("List: Current rune: %c", reader.current)
	l := FormList{}
	for {
		i := reader.read()
		// log.Printf("List: Current form: %T", i)

		r := reader.current
		switch {
		case i == nil && r == ')':
			return l
		case r == eof:
			return reader.errorf("Unclosed list")
		default:
			l.Items = append(l.Items, i)
		}
	}
}

func (reader *Reader) errorf(format string, args ...interface{}) FormError {
	return FormError{
		Val:  reader.input[reader.start:reader.pos],
		Desc: fmt.Sprintf(format, args),
	}
}

func (l *Reader) next() bool {
	if l.pos >= len(l.input) {
		l.width = 0
		l.current = eof
		return false
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	l.current = r
	return true
}

func (l *Reader) ignore() {
	l.start = l.pos
}
