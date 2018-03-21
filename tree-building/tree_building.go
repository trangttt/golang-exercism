package tree

import (
    //"github.com/golang-collections/collections/stack"
    "errors"
    //"fmt"
    "sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}


func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
    number := len(records)
    parents := make(map[int]int)
    for _, record := range records {
        if record.ID == 0 && record.Parent != 0 {
            return nil, errors.New("Incorrect root")
        }
        if record.Parent == record.ID && record.ID != 0 {
            return nil, errors.New("Direct loop")
        }
        if record.Parent > record.ID  {
            return nil, errors.New("Parent is lower then Child")
        }
        if record.ID >= number {
            return nil, errors.New("Non-continuous records")
        }
        _, exist := parents[record.ID]
        if exist {
            return nil, errors.New("Duplicate records")
        } else {
            parents[record.ID] = record.Parent
        }
    }
    children := make([][]int, number)
    for c, p := range parents {
        if c == 0 {
            continue
        }
        children[p] = append(children[p], c)
    }

    // NOTE: till this step, every child is guaranteed to have exactly one parent
    // with ID larger than its ID. meaning the tree is guaranteed with no loop
    // meaning, it is already sorted topologically

    done := make(map[int]Node)
    for i := number - 1; i >=0 ; i-- {
        node := Node{ID: i}
        if len(children[i]) > 0 {
            node.Children = make([]*Node, len(children[i]))
            sort.Ints(children[i])
            for j, c := range children[i] {
                cn := done[c]
                node.Children[j] = &cn
            }
        }
        done[i] = node
    }
    root := done[0]
    return &root, nil
}
