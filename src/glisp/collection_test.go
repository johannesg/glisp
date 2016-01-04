package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_collections(t *testing.T) {
	SkipConvey("Lists", t, func() {
		e := NewEnvironment(nil)

		Convey("list", func() {
			res, err := e.Eval("(list :a 2 \"bb\")")

			So(err, ShouldBeNil)
			expected, err := e.Eval("'(:a 2 \"bb\")")
			So(res, ShouldResemble, expected)
		})

		Convey("first", func() {
			res, err := e.Eval("(first '(:a 2 \"bb\"))")

			So(err, ShouldBeNil)
			expected, err := e.Eval(":a")
			So(res, ShouldResemble, expected)
		})

		Convey("rest", func() {
			res, err := e.Eval("(rest '(:a 2 \"bb\"))")

			So(err, ShouldBeNil)
			expected, err := e.Eval("(2 \"bb\")")
			So(res, ShouldResemble, expected)
		})
	})
}
