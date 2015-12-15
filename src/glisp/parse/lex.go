package parse

type stateFn func(*lexer) stateFn

type lexer struct {
	input string
	start int
	pos   int
	width int
	items chan item
}

func lex(input string) (l *lexer) {
	l = &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()

	return
}

func (l *lexer) run() {
	state := startState
	for state != nil {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(i itemType) {
	l.items <- item{
		typ: i,
		val: l.input[l.start:l.pos],
	}
	l.start = l.pos
}
