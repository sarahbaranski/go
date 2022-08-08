package wordy

import (
	"strconv"
	"strings"
)

type Operation struct {
	number   int
	operator string
}

var operatorsLookup = []string{
	"plus", "minus", "multiplied", "divided", "cubed",
}

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question, "?", "")
	splits := strings.Fields(question)
	var operations []Operation
	current := Operation{0, ""}
	operatorFlag := false
	for i, s := range splits {
		n, _ := strconv.Atoi(s)
		if n != 0 {
			if current.number != 0 {
				return 0, false
			}
			current.number = n
			operatorFlag = false
		} else {
			for _, o := range operatorsLookup {
				if s == o {
					if operatorFlag {
						return 0, false
					}
					current.operator = s
					operatorFlag = true
				}
			}
		}
		if current.operator != "" || i == len(splits)-1 {
			operations = append(operations, current)
			current.number = 0
			current.operator = ""
		}
	}
	if len(operations) == 0 {
		return 0, false
	}
	if len(operations) == 1 {
		if operations[0].operator != "" || operations[0].number == 0 {
			return 0, false
		}
		return operations[0].number, true
	}
	current = operations[0]
	for i := 0; i < len(operations); i++ {
		if current.operator == "" {
			break
		}
		next := Operation{}
		if i+1 < len(operations) {
			next = operations[i+1]
		}
		switch current.operator {
		case operatorsLookup[0]:
			current.number += next.number
		case operatorsLookup[1]:
			current.number -= next.number
		case operatorsLookup[2]:
			current.number *= next.number
		case operatorsLookup[3]:
			current.number /= next.number
		}
		current.operator = next.operator
	}
	return current.number, true
}
