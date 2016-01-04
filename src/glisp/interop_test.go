package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_interop(t *testing.T) {
	Convey("Interop", t, func() {
		e := NewEnvironment(nil)

		Convey("Can call", func() {
			res, err := e.Eval("(.First '(a b c))")

			So(err, ShouldBeNil)

			expected, _ := e.Eval("'a")
			So(res, ShouldResemble, expected)
		})
	})
}
