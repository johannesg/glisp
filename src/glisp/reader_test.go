package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_reader(t *testing.T) {
	Convey("Can parse", t, func() {
		Convey("Errors", func() {
			_, err := NewReader(")").Read()
			So(err, ShouldNotBeNil)
		})

		Convey("Read until end", func() {
			f, err := NewReader("").Read()
			So(f, ShouldBeNil)
			So(err, ShouldBeNil)

			r := NewReader("aaa")
			r.Read()
			f, err = r.Read()
			So(f, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Success", func() {
			f, err := NewReader("aaa").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, Symbol{Name: "aaa"})

			f, err = NewReader("(defn apa)").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				Items: []Form{
					Symbol{Name: "defn"},
					Symbol{Name: "apa"},
				},
			})

			f, err = NewReader("(add 1 2)").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				Items: []Form{
					Symbol{Name: "add"},
					Number{Val: 1},
					Number{Val: 2},
				},
			})

			f, err = NewReader("(print \"Some string\")").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				Items: []Form{
					Symbol{Name: "print"},
					Literal{Val: "Some string"},
				},
			})

			f, err = NewReader("' (add 1 3)").Read()
			So(err, ShouldBeNil)
			So(f, ShouldResemble, &QForm{
				Form: &List{
					Items: []Form{
						Symbol{Name: "add"},
						Number{Val: 1},
						Number{Val: 3},
					},
				}})

			f, err = NewReader("true").Read()
			So(err, ShouldBeNil)
			So(f, ShouldResemble, Boolean(true))

			f, err = NewReader("false").Read()
			So(err, ShouldBeNil)
			So(f, ShouldResemble, Boolean(false))

			f, err = NewReader("[a b]").Read()
			So(err, ShouldBeNil)
			So(f, ShouldResemble, &Vector{
				Items: []Form{
					Symbol{Name: "a"},
					Symbol{Name: "b"},
				},
			})

			f, err = NewReader(":b").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, Keyword{Name: ":b"})
		})
	})
}
