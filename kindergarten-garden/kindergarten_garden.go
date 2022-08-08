package kindergarten

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

// Define the Garden type here.
type Garden struct {
	garden map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(children) == 0 {
		return nil, errors.New("empty children")
	}
	newChildren := append([]string{}, children...)
	sort.Strings(newChildren)
	trimdiagram := strings.TrimLeft(diagram, "\n")
	if trimdiagram == diagram {
		return nil, errors.New("wrong format")
	}
	if IsLower(trimdiagram) {
		return nil, errors.New("wrong diagram")
	}
	newDiagram := strings.Split(trimdiagram, "\n")
	if len(newDiagram[0]) != len(newDiagram[1]) || len(newDiagram[0])%2 != 0 || len(newDiagram[1])%2 != 0 {
		return nil, errors.New("diagrams not the same length")
	}
	row1 := []rune(newDiagram[0])
	row2 := []rune(newDiagram[1])
	if len(newDiagram) == 0 {
		return nil, errors.New("this be empty")
	}
	var g Garden
	g.garden = make(map[string][]string, len(newChildren))
	index := 0
	previousChild := ""
	for _, child := range newChildren {
		if child == previousChild {
			return nil, errors.New("this child repeats")
		} else {
			previousChild = child
		}
		plants := make([]string, 4)
		plants[0] = string(row1[index])
		plants[1] = string(row1[index+1])
		plants[2] = string(row2[index])
		plants[3] = string(row2[index+1])
		index += 2
		g.garden[child] = plants
	}
	return &g, nil
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if plants, ok := g.garden[child]; ok {
		var plantNames []string
		for _, plant := range plants {
			if plant == "R" {
				plantNames = append(plantNames, "radishes")
			} else if plant == "V" {
				plantNames = append(plantNames, "violets")
			} else if plant == "C" {
				plantNames = append(plantNames, "clover")
			} else {
				plantNames = append(plantNames, "grass")
			}
		}
		return plantNames, true
	}
	return nil, false
}
