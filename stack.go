package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Value int
}

type Stack struct {
	Nodes []*Node
	Count int
}

func NewNode(val int) *Node {
	return &Node{val}
}

func NewStack() *Stack {
	return &Stack{}
}

func (st *Stack) Push(val int) {
	st.Nodes = append(st.Nodes, NewNode(val))
	st.Count = len(st.Nodes)
}

func (st *Stack) Pop() (ret int) {
	ret = st.Nodes[max(0, st.Count-1)].Value

	st.Nodes = st.Nodes[:max(0, st.Count-1)]
	st.Count = len(st.Nodes)

	return
}

func (st *Stack) Split() (ret *Stack) {
	ret = NewStack()

	values_to_take := st.Pop()

	for i := 0; i < values_to_take; i++ {
		ret.Push(st.Pop())
	}

	ret.Reverse()

	return
}

func (st *Stack) Reverse() {
	var new []*Node
	size := st.Count

	for i := 0; i < size; i++ {
		new = append(new, NewNode(st.Pop()))
	}

	st.Nodes = new
	st.Count = len(st.Nodes)
}

func (st *Stack) Join(sub *Stack) {
	// Reverse the stack in place since we're pushing values
	sub.Reverse()

	size := sub.Count

	for i := 0; i < size; i++ {
		st.Push(sub.Pop())
	}
}
