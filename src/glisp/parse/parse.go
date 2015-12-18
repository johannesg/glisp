package parse

type ast struct {
	lexer *lexer
}

type parseStateFn func(a *ast) parseStateFn

func parse(input string) (a *ast) {
	l := lex(input)

	a = &ast{lexer: l}

	go a.run()

	return
}

func (a *ast) run() {
	state := parseProgram
	for state != nil {
		state = state(a)
	}
}

func parseProgram(a *ast) parseStateFn {
	return nil
}
