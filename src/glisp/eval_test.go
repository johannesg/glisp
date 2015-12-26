package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_eval(t *testing.T) {
	Convey("Evaluation", t, func() {
		Convey("Basic", func() {
			f, _ := NewReader("4").Read()
			res, err := f.Eval()

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 4})
		})

		Convey("Addition", func() {
			f, _ := NewReader("(add 1 4)").Read()
			res, err := f.Eval()

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 5})

			f, _ = NewReader("(add 1 (add 9 4))").Read()
			res, err = f.Eval()

			So(err, ShouldBeNil)
			So(res, ShouldResemble, Number{Val: 14})
		})
	})
}
