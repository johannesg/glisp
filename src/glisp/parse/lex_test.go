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
		l := lex(" ")

		i := <-l.items

		So(i.typ, ShouldEqual, itemEOF)
	})

	Convey("Should parse left paren", t, func() {
		l := lex(" (")

		i := <-l.items

		So(i.typ, ShouldEqual, itemLeftParen)
		So(i.val, ShouldEqual, "(")
	})
}
