package reader

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_lexer(t *testing.T) {
	Convey("Should parse empty string", t, func() {
		l := lex("")

		i := <-l.tokens

		So(i.typ, ShouldEqual, tokenEOF)
	})

	Convey("Should parse whitespace", t, func() {
		l := lex(" \n \t ")

		i := <-l.tokens

		So(i.typ, ShouldEqual, tokenEOF)
	})

	Convey("Should parse parens", t, func() {
		Convey("Left paren", func() {
			l := lex("(")
			VerifyNextToken(l, tokenDelim, "(")
			VerifyNextToken(l, tokenEOF, "")
		})
		Convey("Right paren", func() {
			l := lex(")")
			VerifyNextToken(l, tokenDelim, ")")
			VerifyNextToken(l, tokenEOF, "")
		})
		Convey("Both", func() {
			l := lex(" () ( ) ")
			VerifyNextToken(l, tokenDelim, "(")
			VerifyNextToken(l, tokenDelim, ")")
			VerifyNextToken(l, tokenDelim, "(")
			VerifyNextToken(l, tokenDelim, ")")
			VerifyNextToken(l, tokenEOF, "")
		})
		// So(i.typ, ShouldEqual, tokenLeftParen)
		// So(i.val, ShouldEqual, "(")
	})

	Convey("Identifiers", t, func() {
		Convey("Alphanumeric", func() {
			l := lex("_abc123")
			VerifyNextToken(l, tokenIdentifier, "_abc123")
			VerifyNextToken(l, tokenEOF, "")
		})

		Convey("Multiple identifiers", func() {
			l := lex(" _abc123 bbb\nccc")
			VerifyNextToken(l, tokenIdentifier, "_abc123")
			VerifyNextToken(l, tokenIdentifier, "bbb")
			VerifyNextToken(l, tokenIdentifier, "ccc")
			VerifyNextToken(l, tokenEOF, "")
		})
	})

	Convey("Numbers", t, func() {
		l := lex("1234567890")
		VerifyNextToken(l, tokenNumber, "1234567890")
		VerifyNextToken(l, tokenEOF, "")

		l = lex("-1234567890")
		VerifyNextToken(l, tokenNumber, "-1234567890")
		VerifyNextToken(l, tokenEOF, "")

		l = lex("-123456.7890")
		VerifyNextToken(l, tokenNumber, "-123456.7890")
		VerifyNextToken(l, tokenEOF, "")

		l = lex("-123456.78.90")
		VerifyError(l)
	})

	Convey("Strings", t, func() {
		l := lex("  \"a nice string, 11334.9 ;'[][\" ")
		VerifyNextToken(l, tokenString, "a nice string, 11334.9 ;'[][")

		l = lex("(\"A string\")")
		VerifyNextToken(l, tokenDelim, "(")
		VerifyNextToken(l, tokenString, "A string")
		VerifyNextToken(l, tokenDelim, ")")
	})

	Convey("Misc", t, func() {
		l := lex(`(defn foo [a b] 
  (add a b))
`)

		VerifyNextToken(l, tokenDelim, "(")
		VerifyNextToken(l, tokenIdentifier, "defn")
		VerifyNextToken(l, tokenIdentifier, "foo")
		VerifyNextToken(l, tokenDelim, "[")
		VerifyNextToken(l, tokenIdentifier, "a")
		VerifyNextToken(l, tokenIdentifier, "b")
		VerifyNextToken(l, tokenDelim, "]")
		VerifyNextToken(l, tokenDelim, "(")
		VerifyNextToken(l, tokenIdentifier, "add")
		VerifyNextToken(l, tokenIdentifier, "a")
		VerifyNextToken(l, tokenIdentifier, "b")
		VerifyNextToken(l, tokenDelim, ")")
		VerifyNextToken(l, tokenDelim, ")")
		VerifyNextToken(l, tokenEOF, "")
	})

	Convey("Quote", t, func() {
		l := lex("(defn 'add (a b))")

		VerifyNextToken(l, tokenDelim, "(")
		VerifyNextToken(l, tokenIdentifier, "defn")
		VerifyNextToken(l, tokenQuote, "'")
		VerifyNextToken(l, tokenIdentifier, "add")
		VerifyNextToken(l, tokenDelim, "(")
		VerifyNextToken(l, tokenIdentifier, "a")
		VerifyNextToken(l, tokenIdentifier, "b")
		VerifyNextToken(l, tokenDelim, ")")
	})

	Convey("Comments", t, func() {
		l := lex(`2 ; the number 2
; a comment of no interest
5; the number 5`)

		VerifyNextToken(l, tokenNumber, "2")
		VerifyNextToken(l, tokenComment, " the number 2")
		VerifyNextToken(l, tokenComment, " a comment of no interest")
		VerifyNextToken(l, tokenNumber, "5")
		VerifyNextToken(l, tokenComment, " the number 5")
	})
}

func VerifyError(l *lexer) {
	i := <-l.tokens
	So(i.typ, ShouldEqual, tokenError)
}

func VerifyNextToken(l *lexer, t tokenType, v string) {
	i := <-l.tokens
	So(i.typ, ShouldEqual, t)
	So(i.val, ShouldEqual, v)
}
