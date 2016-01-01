package main

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
		case strings.ContainsRune("+-", r):
			return lexIdentifierOrNumber
		case strings.ContainsRune("=", r), isAlpha(r):
			return lexIdentifier
		case r == '-', unicode.IsNumber(r):
			return lexNumber
		case r == '"':
			return lexString
		case r == '\'':
			l.emit(tokenQuote)
			return lexProgram
		case r == ';':
			return lexComment
		default:
			return l.errorf("Unknown token: %s", l.input[l.start:l.pos])
		}
	}
	l.emit(tokenEOF)
	return nil
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == eof
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsNumber(r)
}

func isMacro(r rune) bool {
	return strings.ContainsRune("\";'@^`~()[]{}\\%#", r)
}

func isTerminatingMacro(r rune) bool {
	return r != '#' && r != '\'' && r != '%' && isMacro(r)
}

func lexIdentifierOrNumber(l *lexer) stateFn {
	r := l.next()
	if unicode.IsNumber(r) {
		return lexNumber
	} else {
		l.backup()
		return lexIdentifier
	}
}

func lexIdentifier(l *lexer) stateFn {
	for {
		r := l.next()
		if isWhitespace(r) || isTerminatingMacro(r) {
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

		case isWhitespace(r) || isMacro(r):
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

func lexComment(l *lexer) stateFn {
	l.ignore()
	for {
		switch l.next() {
		case eof, '\n':
			l.backup()
			l.emit(tokenComment)
			return lexProgram
		}
	}
}
