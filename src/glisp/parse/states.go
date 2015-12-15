package parse

func startState(l *lexer) stateFn {
	l.emit(itemEOF)
	return nil
}
