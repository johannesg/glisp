package parse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_lexer(t *testing.T) {
	Convey("Should parse empty string", t, func() {
		l := lex("")

		i := <-l.items

		So(i.typ, ShouldEqual, itemEOF)
	})

	Convey("Should parse whitespace", t, func() {
		l := lex(" \n \t ")

		i := <-l.items

		So(i.typ, ShouldEqual, itemEOF)
	})

	Convey("Should parse parens", t, func() {
		Convey("Left paren", func() {
			l := lex("(")
			VerifyNextToken(l, itemDelim, "(")
			VerifyNextToken(l, itemEOF, "")
		})
		Convey("Right paren", func() {
			l := lex(")")
			VerifyNextToken(l, itemDelim, ")")
			VerifyNextToken(l, itemEOF, "")
		})
		Convey("Both", func() {
			l := lex(" () ( ) ")
			VerifyNextToken(l, itemDelim, "(")
			VerifyNextToken(l, itemDelim, ")")
			VerifyNextToken(l, itemDelim, "(")
			VerifyNextToken(l, itemDelim, ")")
			VerifyNextToken(l, itemEOF, "")
		})
		// So(i.typ, ShouldEqual, itemLeftParen)
		// So(i.val, ShouldEqual, "(")
	})

	Convey("Identifiers", t, func() {
		Convey("Alphanumeric", func() {
			l := lex("_abc123")
			VerifyNextToken(l, itemIdentifier, "_abc123")
			VerifyNextToken(l, itemEOF, "")
		})

		Convey("Multiple identifiers", func() {
			l := lex(" _abc123 bbb\nccc")
			VerifyNextToken(l, itemIdentifier, "_abc123")
			VerifyNextToken(l, itemIdentifier, "bbb")
			VerifyNextToken(l, itemIdentifier, "ccc")
			VerifyNextToken(l, itemEOF, "")
		})
	})

	Convey("Numbers", t, func() {
		l := lex("1234567890")
		VerifyNextToken(l, itemNumber, "1234567890")
		VerifyNextToken(l, itemEOF, "")

		l = lex("-1234567890")
		VerifyNextToken(l, itemNumber, "-1234567890")
		VerifyNextToken(l, itemEOF, "")

		l = lex("-123456.7890")
		VerifyNextToken(l, itemNumber, "-123456.7890")
		VerifyNextToken(l, itemEOF, "")

		l = lex("-123456.78.90")
		VerifyError(l)
	})

	Convey("Strings", t, func() {
		l := lex("  \"a nice string, 11334.9 ;'[][\" ")
		VerifyNextToken(l, itemString, "a nice string, 11334.9 ;'[][")
	})

	Convey("Misc", t, func() {
		l := lex(`(defn foo [a b] 
  (add a b))
`)

		VerifyNextToken(l, itemDelim, "(")
		VerifyNextToken(l, itemIdentifier, "defn")
		VerifyNextToken(l, itemIdentifier, "foo")
		VerifyNextToken(l, itemDelim, "[")
		VerifyNextToken(l, itemIdentifier, "a")
		VerifyNextToken(l, itemIdentifier, "b")
		VerifyNextToken(l, itemDelim, "]")
		VerifyNextToken(l, itemDelim, "(")
		VerifyNextToken(l, itemIdentifier, "add")
		VerifyNextToken(l, itemIdentifier, "a")
		VerifyNextToken(l, itemIdentifier, "b")
		VerifyNextToken(l, itemDelim, ")")
		VerifyNextToken(l, itemDelim, ")")
		VerifyNextToken(l, itemEOF, "")
	})
}

func VerifyError(l *lexer) {
	i := <-l.items
	So(i.typ, ShouldEqual, itemError)
}

func VerifyNextToken(l *lexer, t itemType, v string) {
	i := <-l.items
	So(i.typ, ShouldEqual, t)
	So(i.val, ShouldEqual, v)
}
