package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Koan(t *testing.T) {
	return

	Convey("Koan 01", t, func() {
		assertTrue(
			"We shall contemplate truth by testing reality, via equality",
			"(= true true)")

		assertTrue(
			"To understand reality, we must compare our expectations against reality",
			"(= 2 (+ 1 1))")

		assertTrue(
			"You can test equality of many things",
			"(= (+ 3 4) 7 (+ 2 5))")

		assertTrue(
			"Some things may appear different, but be the same",
			"(= true (= 2 2/1))")

		assertTrue(
			"You cannot generally float to heavens of integers",
			"(= false (= 2 2.0))")

		assertTrue(
			"But a looser equality is also possible",
			"(= true (== 2.0 2))")

		assertTrue(
			"Something is not equal to nothing",
			"(= true (not (= 1 nil)))")

		assertTrue(
			"Strings, and keywords, and symbols: oh my!",
			`(= false (= "hello" :hello 'hello))`)

		assertTrue(
			"Make a keyword with your keyboard",
			`(= :hello (keyword "hello"))`)

		assertTrue(
			"Symbolism is all around us",
			`(= 'hello (symbol "hello"))`)

		assertTrue(
			"When things cannot be equal, they must be different",
			"(not= :fill-in-the-blank __))")

	})
}

func assertTrue(desc string, lisp string) {
	Convey(desc, func() {
		e := NewEnvironment(nil)
		f, err := NewReader(lisp).Read()

		So(err, ShouldBeNil)

		r, err := f.Eval(e)

		So(err, ShouldBeNil)
		So(r, ShouldResemble, Boolean(true))
	})
}
