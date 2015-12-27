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

		Convey("Eval", func() {
			res, err := e.Eval("(eval '(+ 2 3))")
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
				res, err := e.Eval("(fn [x y] '(+ x y))")

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

			Convey("Can call", func() {
				res, err := e.Eval("( (fn [x y] '(+ x y)) 3 5)")

				So(err, ShouldBeNil)
				So(res, ShouldResemble, Number{Val: 8})
			})

			Convey("Curry", func() {
				res, err := e.Eval("( (fn [x y] '(+ x y)) 3)")

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
	})
}
