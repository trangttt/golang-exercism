package linkedlist

import "errors"

type Element struct {
    data int
    next *Element
}
type List struct {
    head *Element
    size int
}

func New(input []int) *List {
    if len(input) == 0 {
        return &List{head: nil,
                    size: 0}
    }
    head := Element{data: input[0],
                    next: nil}
	current := &head
    for _, elem := range input[1:] {
        current.next = &Element{data: elem,
                                next: nil}
		current = current.next
    }
    return &List{head: &head, size: len(input)}
}

func (l *List) Size() int {
     return l.size
 }

func (l *List) Push(v int) {
	if l.size == 0 {
	    l.head = &Element{data: v, next: nil}
	    l.size++
	    return
    }
    current := l.head
    for current.next != nil {
       current = current.next
    }
    current.next = &Element{data: v, next: nil}
    l.size++
}
func (l *List) Pop() (int, error) {
	if l.size == 0 {
	    return -1, errors.New("Empty List")
    }
    current := l.head
    prev := &Element{}

    for current.next != nil {
    	prev = current
        current = current.next
    }

    if prev != nil {
        prev.next = nil
    }

    l.size--
    if l.size == 0 { l.head = nil}
    return current.data, nil
}

func (l *List) Array() []int {
    ret := make([]int, l.size)
    current := l.head
    index := 0
    for current != nil {
       ret[index] = current.data
       index++
       current = current.next
    }
    return ret
}

func (l *List) Reverse() *List {
	arr := make([]int, l.size)
	index := l.size - 1
	current := l.head
	for index >=0 {
	    arr[index] = current.data
	    current = current.next
	    index--
    }
    return New(arr)
}



