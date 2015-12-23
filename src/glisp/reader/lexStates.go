package parse

import (
	"strings"
	"unicode"
)

func lexProgram(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			break
		}
		switch {
		case unicode.IsSpace(r):
			l.ignore()
		case strings.ContainsRune("()[]", r):
			l.emit(tokenDelim)
		case isAlpha(r):
			return lexIdentifier
		case r == '-', unicode.IsNumber(r):
			return lexNumber
		case r == '"':
			return lexString
		default:
			return l.errorf("Unknown token: %v", r)
		}
	}
	l.emit(tokenEOF)
	return nil
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n'
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsNumber(r)
}

func lexIdentifier(l *lexer) stateFn {
	for {
		r := l.next()
		if !isAlphaNumeric(r) {
			l.backup()
			l.emit(tokenIdentifier)
			break
		}
	}
	return lexProgram
}

func lexNumber(l *lexer) stateFn {
	foundComma := false
	for {
		r := l.next()
		switch {
		case r == '.':
			if foundComma {
				return l.errorf("Tried to parse a number but found 2 commas")
			}
			foundComma = true

		case !unicode.IsNumber(r):
			l.backup()
			l.emit(tokenNumber)
			return lexProgram
		}
	}
}

func lexString(l *lexer) stateFn {
	l.ignore()
Loop:
	for {
		switch l.next() {
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.backup()
	l.emit(tokenString)
	l.next()
	l.ignore()
	return lexProgram
}
