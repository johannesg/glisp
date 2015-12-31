package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_eval(t *testing.T) {
	Convey("Evaluation", t, func() {
		e := NewEnvironment(nil)
		Convey("Basic", func() {
			f, _ := NewReader("4").Read()
			res, err := f.Eval(e)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 4})
		})

		Convey("Do", func() {
			res, err := e.Eval("(do '(+ 2 3))")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 5})
		})

		Convey("Addition", func() {
			f, _ := NewReader("(add 1 4)").Read()
			res, err := f.Eval(e)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 5})

			f, _ = NewReader("(add 1 (add 9 4))").Read()
			res, err = f.Eval(e)

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 14})
		})

		Convey("Functions", func() {
			Convey("Can define", func() {
				res, err := e.Eval("(fn [x y] (+ x y))")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, UserFunction{
					Args: []Symbol{
						Symbol{Name: "x"},
						Symbol{Name: "y"},
					},
					Body: &List{
						Items: []Form{
							Symbol{Name: "+"},
							Symbol{Name: "x"},
							Symbol{Name: "y"},
						}},
				})
			})

			Convey("Using defn", func() {
				res, err := e.Eval("(defn add2 [x y] (+ x y))")

				So(err, ShouldBeNil)

				res, err = e.Eval("(add2 3 5)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, Number{Val: 8})
			})

			Convey("Can call", func() {
				res, err := e.Eval("( (fn [x y] (+ x y)) 3 5)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, Number{Val: 8})
			})

			SkipConvey("Curry", func() {
				res, err := e.Eval("( (fn [x y] (+ x y)) 3)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, Number{Val: 8})

				So(res, ShouldResemble, UserFunction{
					Args: []Symbol{
						Symbol{Name: "y"},
					},
					Body: &List{
						Items: []Form{
							Symbol{Name: "+"},
							Number{Val: 3},
							Symbol{Name: "y"},
						}},
				})
			})
		})

		Convey("Macros", func() {
			Convey("Can define", func() {
				res, err := e.Eval("(defmacro add2 [x y] (+ x y))")

				So(err, ShouldBeNil)

				res, err = e.Eval("(add2 3 5)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, Number{Val: 8})
			})

			Convey("Can expand", func() {
				res, err := e.Eval("(defmacro add2 [x y] (+ x y))")

				So(err, ShouldBeNil)

				res, err = e.Eval("(macroexpand add2 2 4)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, &List{
					Items: []Form{
						Symbol{Name: "+"},
						Number{Val: 2},
						Number{Val: 4},
					},
				})
			})
		})
	})
}
