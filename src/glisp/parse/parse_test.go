package parse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_parse(t *testing.T) {
	Convey("Can parse", t, func() {

		parse(" (defn apa)")

	})
}
