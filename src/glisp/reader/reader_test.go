package reader

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
			So(f, ShouldResemble, Symbol{name: "aaa"})

			f, _ = NewReader("(defn apa)").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				items: []Form{
					Symbol{name: "defn"},
					Symbol{name: "apa"},
				},
			})

			f, _ = NewReader("(add 1 2)").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				items: []Form{
					Symbol{name: "add"},
					Number{val: 1},
					Number{val: 2},
				},
			})

			f, _ = NewReader("(print \"Some string\")").Read()

			So(err, ShouldBeNil)
			So(f, ShouldResemble, &List{
				items: []Form{
					Symbol{name: "print"},
					Literal{val: "Some string"},
				},
			})
		})

	})
}
