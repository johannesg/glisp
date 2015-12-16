package parse

func lexProgram(l *lexer) stateFn {
	r := l.peek()
	for r != eof {
		switch {
		case isWhitespace(r):
			l.ignore()
		case r == '(':
			l.emit(itemLeftParen)
		case r == ')':
			l.emit(itemRightParen)
		}
		r = l.next()
	}
	l.emit(itemEOF)
	return nil
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n'
}
