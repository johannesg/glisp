package parse

type Form interface {
}

type List struct {
	items []Form
}

type Symbol struct {
	name string
}

type Literal struct {
	val string
}

type Number struct {
	val int
}
