package parse

import "fmt"

type itemType int

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemError:
		return i.val
	case itemEOF:
		return "EOF"
	}

	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

const (
	itemError = iota
	itemEOF
	itemLeftParen
	itemRightParen
	itemNumber
	itemString
)
