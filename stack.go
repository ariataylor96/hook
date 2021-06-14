package main

type Stack struct {
	Values []int
}

func (st *Stack) Push(val int) {
	st.Values = append(st.Values, val)
}

func (st *Stack) Pop() (ret int) {
	size := len(st.Values)

	ret = st.Values[size-1]
	st.Values = st.Values[:size]

	return
}

func (st *Stack) Split() (ret Stack) {
	values_to_take := st.Pop()
	size := len(st.Values)

	ret = Stack{st.Values[size-values_to_take : size]}
	st.Values = st.Values[0 : size-values_to_take]

	return
}

func (st *Stack) Reverse() {
	var new []int

	// This gets modified as we iterate
	size := len(st.Values)

	for i := 0; i < size; i++ {
		new = append(new, st.Pop())
	}

	st.Values = new
}

func (st *Stack) Join(sub *Stack) {
	// Reverse the stack in place since we're pushing values
	sub.Reverse()

	for _, val := range sub.Values {
		st.Push(val)
	}
}
