package parse

import "unicode"

func lexProgram(l *lexer) stateFn {
	r := l.peek()
	for r != eof {
		r = l.next()

		switch {
		case unicode.IsSpace(r):
			l.ignore()
		case r == '(':
			l.emit(itemLeftParen)
		case r == ')':
			l.emit(itemRightParen)
		case isAlpha(r):
			return lexIdentifier
		case r == '-', unicode.IsNumber(r):
			return lexNumber
		}
	}
	l.emit(itemEOF)
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
			l.emit(itemIdentifier)
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
			l.emit(itemNumber)
			return lexProgram
		}
	}

}
