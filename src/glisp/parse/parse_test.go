package parse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_parse(t *testing.T) {
	Convey("Can parse", t, func() {
		Convey("Errors", func() {
		})

		Convey("Success", func() {
			a := parse("aaa")

			So(a.nodes, ShouldHaveLength, 1)
			So(a.nodes[0], ShouldHaveSameTypeAs, astSymbol{})

			a = parse("(defn apa)")

			So(a.nodes, ShouldHaveLength, 1)
			So(a.nodes[0], ShouldHaveSameTypeAs, astList{})
			So(a.nodes[0].(astList).items, ShouldHaveLength, 2)

			a = parse("(add 1 2)")

			So(a.nodes, ShouldHaveLength, 1)
			So(a.nodes[0], ShouldHaveSameTypeAs, astList{})
			So(a.nodes[0].(astList).items, ShouldHaveLength, 3)
			So(a.nodes[0].(astList).items[0], ShouldHaveSameTypeAs, astSymbol{})
			So(a.nodes[0].(astList).items[1], ShouldHaveSameTypeAs, astLiteral{})
			So(a.nodes[0].(astList).items[2], ShouldHaveSameTypeAs, astLiteral{})

			a = parse("(print \"Some string\")")

			So(a.nodes, ShouldHaveLength, 1)
			So(a.nodes[0], ShouldHaveSameTypeAs, astList{})
			So(a.nodes[0].(astList).items, ShouldHaveLength, 2)
			So(a.nodes[0].(astList).items[0], ShouldHaveSameTypeAs, astSymbol{})
			So(a.nodes[0].(astList).items[1], ShouldHaveSameTypeAs, astLiteral{})
			So(a.nodes[0].(astList).items[1].(astLiteral).val, ShouldEqual, "Some string")
		})

	})
}
