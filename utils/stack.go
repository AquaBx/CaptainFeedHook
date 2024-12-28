package utils

type Stack[T any] []T

func (p *Stack[T]) Pop() T {
	if len(*p) > 0 {
		res := (*p)[len(*p)-1]
		*p = (*p)[0 : len(*p)-1]
		return res
	} else {
		panic("Empty stack")
	}
}

func (p *Stack[T]) Push(el T) {
	*p = append(*p, el)
}

func (p *Stack[T]) Length() int {
	return len(*p)
}
