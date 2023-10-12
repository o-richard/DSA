package tree

type rbNode struct {
	data int
	isRed bool
	left, right, parent *rbNode
}

type rb struct {
	root *rbNode
}

func InitRB() rb {
	return rb{}
}