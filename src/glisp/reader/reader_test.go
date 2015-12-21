package reader

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_reader(t *testing.T) {
	Convey("Empty", t, func() {
		r := New("")

		So(r.Read(), ShouldHaveSameTypeAs, FormEmpty{})

		r = New("   ")
		So(r.Read(), ShouldHaveSameTypeAs, FormEmpty{})
	})

	Convey("Symbols", t, func() {
		r := New("aaa")
		f := r.Read()

		So(f, ShouldHaveSameTypeAs, FormSymbol{})
		So(f.(FormSymbol).Name, ShouldEqual, "aaa")

		r = New("  aaa  ")
		f = r.Read()

		So(f, ShouldHaveSameTypeAs, FormSymbol{})
		So(f.(FormSymbol).Name, ShouldEqual, "aaa")

		r = New("  +a-b?  ")
		f = r.Read()

		So(f, ShouldHaveSameTypeAs, FormSymbol{})
		So(f.(FormSymbol).Name, ShouldEqual, "+a-b?")
	})

	Convey("Numbers", t, func() {
		r := New("  1234  ")
		f := r.Read()

		So(f, ShouldHaveSameTypeAs, FormNumber{})
		So(f.(FormNumber).Val, ShouldEqual, 1234)

		r = New("  -1234  ")
		f = r.Read()

		So(f, ShouldHaveSameTypeAs, FormNumber{})
		So(f.(FormNumber).Val, ShouldEqual, -1234)

		r = New("  +1234  ")
		f = r.Read()

		So(f, ShouldHaveSameTypeAs, FormNumber{})
		So(f.(FormNumber).Val, ShouldEqual, +1234)
	})
}
