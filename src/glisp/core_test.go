package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_core(t *testing.T) {
	Convey("Lists", t, func() {
		e := NewEnvironment(nil)
		_, err := e.Load("core.clj")
		So(err, ShouldBeNil)

		SkipConvey("list", func() {
			res, err := e.Eval(`(list :a :b :c)`)
			exp, _ := e.Eval(`'(:a :b :c)`)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, exp)
		})

		Convey("first", func() {
			res, err := e.Eval(`(first '(:a :b :c))`)
			exp, _ := e.Eval(":a")

			So(err, ShouldBeNil)
			So(res, ShouldResemble, exp)
		})

		Convey("rest", func() {
			res, err := e.Eval(`(rest '(:a :b :c))`)
			exp, _ := e.Eval(`'(:b :c)`)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, exp)
		})

		Convey("cons", func() {
			res, err := e.Eval(`(cons :a '(:b :c))`)
			exp, _ := e.Eval(`'(:a :b :c)`)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, exp)
		})

		Convey("conj", func() {
			res, err := e.Eval(`(conj '(:b :c) :a)`)
			exp, _ := e.Eval(`'(:a :b :c)`)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, exp)
		})
	})
}
