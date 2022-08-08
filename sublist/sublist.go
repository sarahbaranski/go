package sublist

type Relation string

func Sublist(l1, l2 []int) string {
	len1 := len(l1)
	len2 := len(l2)
	switch {
	case len1 == 0 && len2 == 0:
		return "equal"
	case len1 == 0:
		return "sublist"
	case len2 == 0:
		return "superlist"
	case len1 > len2:
		if isSublist(l1, l2) {
			return "superlist"
		}
		return "unequal"
	case len1 == len2:
		if isSublist(l1, l2) {
			return "equal"
		}
		return "unequal"
	default:
		if isSublist(l2, l1) {
			return "sublist"
		}
		return "unequal"
	}
}

func isSublist(list, sub []int) bool {
	var sublist bool
	for i := 0; i <= len(list)-len(sub); i++ {
		sublist = true
		for j := 0; j < len(sub); j++ {
			if list[i+j] != sub[j] {
				sublist = false
				break
			}
		}
		if sublist {
			break
		}
	}
	return sublist
}
