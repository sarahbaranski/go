package tree

import (
	"errors"
	"sort"
)

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
	if len(records) == 1 {
		if records[0].ID != 0 {
			return nil, errors.New("root ID must be 0")
		}
		return &Node{ID: records[0].ID}, nil
	}
	// sort records by child IDs
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	// sort records by Parent IDs
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Parent < records[j].Parent
	})
	// check for duplcate nodes and roots in records
	if ContainsDuplicate(records) {
		return nil, errors.New("records cannot contain duplicates")
	}
	// node: the initial and current node
	node := Node{ID: records[0].ID}
	// counts records to verify if all are converted to nodes
	nodeCount := 1
	for i := 1; i < len(records); i++ {
		record := records[i]
		// verify ID is less than parent
		if record.ID < record.Parent {
			return nil, errors.New("ID cannot be less than Parent ID")
		}
		// not really sure why we need...but it fixes the "non-continuous" edge case
		if record.ID >= len(records) {
			return nil, errors.New("eh, id is greater than length, i guess that shouldnt happen")
		}
		// add record as child node if current node.ID is the record.Parent ID
		if node.ID == record.Parent {
			node.Children = append(node.Children, &Node{ID: record.ID})
			nodeCount++
		} else {
			// loop through child nodes to find child with ID for record.Parent
			for j := 0; j < len(node.Children); j++ {
				child := node.Children[j]
				if child.ID == record.Parent {
					child.Children = append(child.Children, &Node{ID: record.ID})
					nodeCount++
				}
			}
		}
	}
	// verify all records were converted to nodes based on nodeCount
	if nodeCount != len(records) {
		return nil, errors.New("diconnected parent nodes")
	}
	return &node, nil
}

/*
	checks records for duplicate entries
*/
func ContainsDuplicate(records []Record) bool {
	previous := records[0]
	for i := 1; i < len(records); i++ {
		if previous.ID == records[i].ID && previous.Parent == records[i].Parent {
			return true
		}
		previous = records[i]
	}
	return false
}
