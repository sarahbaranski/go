package tree

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// root := &Node{ID: 0}
	if len(records) == 1 {
		return &Node{ID: records[0].ID}, nil
	}
	//array we made:
	// n := []*Node{&Node{ID: records[0].ID}}

	//tried making into a map:
	n := map[int]int*Node{&Node{ID: records.ID}}
	for _, v := range records {
		temp := &Node{ID: records.ID}
		if val, ok := records.ID; ok {
			n[records.Parent].Children = append(n[records.Parent].Children, temp)
		}
	}
	//looping through the array
	// for i := 0; i < len(records); i++ {
	// 	temp := &Node{ID: records[i].ID}
	// 	if len(n) == records[i].ID {
	// 		n[records[i].Parent].Children = append(n[records[i].Parent].Children, temp)
	// 	}
	// }
	return root, nil
}
