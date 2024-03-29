package kindergarten

import (
	"reflect"
	"sort"
	"testing"
)

type lookup struct {
	child  string
	plants []string
	ok     bool
}

type gardenTest struct {
	number   int
	diagram  string
	children []string
	ok       bool
	lookups  []lookup
}

var tests = []gardenTest{
	{1, `
RC
GG`, []string{"Alice"}, true, []lookup{
		{"Alice", []string{"radishes", "clover", "grass", "grass"}, true},
	}},
	{2, `
VC
RC`, []string{"Alice"}, true, []lookup{
		{"Alice", []string{"violets", "clover", "radishes", "clover"}, true},
	}},
	{3, `
VVCG
VVRC`, []string{"Alice", "Bob"}, true, []lookup{
		{"Bob", []string{"clover", "grass", "radishes", "clover"}, true},
	}},
	{4, `
VVCCGG
VVCCGG`, []string{"Alice", "Bob", "Charlie"}, true, []lookup{
		{"Bob", []string{"clover", "clover", "clover", "clover"}, true},
		{"Charlie", []string{"grass", "grass", "grass", "grass"}, true},
	}},
	test5, // full garden test
	test6, // out of order names test

	// failure tests
	{7, "RC\nGG", []string{"Alice"}, false, nil}, // wrong diagram format
	{8, `
RCCC
GG`, []string{""}, false, nil}, // mismatched rows
	{9, `
RCC
GGC`, []string{"Alice"}, false, nil}, // odd number of cups
	{10, `
RCCC
GGCC`, []string{"Alice", "Alice"}, false, nil}, // duplicate name
	{11, `
rc
gg`, []string{"Alice"}, false, nil}, // invalid cup codes
	{12, `
RC
GG`, []string{"Alice"}, true, []lookup{ // lookup invalid name
		{"Bob", []string{"radishes", "clover", "grass", "grass"}, false},
	}},
}

// full garden test
var test5 = gardenTest{5, `
VRCGVVRVCGGCCGVRGCVCGCGV
VRCCCGCRRGVCGCRVVCVGCGCV`, []string{
	"Alice", "Bob", "Charlie", "David", "Eve", "Fred",
	"Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"}, true, []lookup{
	{"Alice", []string{"violets", "radishes", "violets", "radishes"}, true},
	{"Bob", []string{"clover", "grass", "clover", "clover"}, true},
	{"Charlie", []string{"violets", "violets", "clover", "grass"}, true},
	{"David", []string{"radishes", "violets", "clover", "radishes"}, true},
	{"Eve", []string{"clover", "grass", "radishes", "grass"}, true},
	{"Fred", []string{"grass", "clover", "violets", "clover"}, true},
	{"Ginny", []string{"clover", "grass", "grass", "clover"}, true},
	{"Harriet", []string{"violets", "radishes", "radishes", "violets"}, true},
	{"Ileana", []string{"grass", "clover", "violets", "clover"}, true},
	{"Joseph", []string{"violets", "clover", "violets", "grass"}, true},
	{"Kincaid", []string{"grass", "clover", "clover", "grass"}, true},
	{"Larry", []string{"grass", "violets", "clover", "violets"}, true},
}}

// out of order names test
var (
	test6names = []string{"Samantha", "Patricia", "Xander", "Roger"}
	test6      = gardenTest{6, `
VCRRGVRG
RVGCCGCV`, append([]string{}, test6names...), true, []lookup{
		{"Patricia", []string{"violets", "clover", "radishes", "violets"}, true},
		{"Roger", []string{"radishes", "radishes", "grass", "clover"}, true},
		{"Samantha", []string{"grass", "violets", "clover", "grass"}, true},
		{"Xander", []string{"radishes", "grass", "clover", "violets"}, true},
	}}
)

func TestGarden(t *testing.T) {
	for _, test := range tests {
		g, err := NewGarden(test.diagram, test.children)
		if test.ok {
			if err != nil {
				t.Fatalf("NewGarden test %d returned error %q.  Error not expected.",
					test.number, err)
			}

			for _, l := range test.lookups {
				switch plants, ok := g.Plants(l.child); {
				case ok != l.ok:
					t.Fatalf("Garden %d lookup %s returned ok = %t, want %t.",
						test.number, l.child, ok, l.ok)
				case ok && !reflect.DeepEqual(plants, l.plants):
					t.Fatalf("Garden %d lookup %s = %q, want %q.",
						test.number, l.child, plants, l.plants)
				}
			}
		} else { // negative tests; expecting error
			// check err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("NewGarden test %d. Expected error but got nil.", test.number)
			}
		}
	}
}

// The lazy way to meet the alphabetizing requirement is with sort.Strings
// on the argument slice.  That's an in-place sort though and it's bad practice
// to have a side effect.
func TestNamesNotModified(t *testing.T) {
	cp := append([]string{}, test6names...)
	_, err := NewGarden(test6.diagram, cp)
	if err != nil || sort.StringsAreSorted(cp) {
		t.Fatalf("error in test setup: TestNamesNotModified requires valid garden and unsorted children")
	}
	if !reflect.DeepEqual(cp, test6names) {
		t.Fatalf("NewGarden modified children argment. Arguments should not be modified.")
	}
}

// A test taken from the Ruby tests.  It checks that Garden objects
// are self-contained and do not rely on package variables.
func TestTwoGardens(t *testing.T) {
	diagram := `
VCRRGVRG
RVGCCGCV`
	g1, err1 := NewGarden(diagram, []string{"Alice", "Bob", "Charlie", "Dan"})
	g2, err2 := NewGarden(diagram, []string{"Bob", "Charlie", "Dan", "Erin"})
	if err1 != nil || err2 != nil {
		t.Fatalf("error in test setup: Two garden test needs valid gardens")
	}
	tf := func(g *Garden, n int, child string, expPlants []string) {
		switch plants, ok := g.Plants(child); {
		case !ok:
			t.Fatalf("error in test setup: Garden %d lookup %s returned ok = false, want true. Check if the child exists in the garden", n, child)
		case !reflect.DeepEqual(plants, expPlants):
			t.Fatalf("Garden %d lookup %s = %q, want %q.",
				n, child, plants, expPlants)
		}
	}
	tf(g1, 1, "Bob", []string{"radishes", "radishes", "grass", "clover"})
	tf(g2, 2, "Bob", []string{"violets", "clover", "radishes", "violets"})
	tf(g1, 1, "Charlie", []string{"grass", "violets", "clover", "grass"})
	tf(g2, 2, "Charlie", []string{"radishes", "radishes", "grass", "clover"})
}

func BenchmarkNewGarden(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			NewGarden(test.diagram, test.children)
		}
	}
}

func BenchmarkGarden_Plants(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	g, err := NewGarden(test5.diagram, test5.children)
	if err != nil {
		b.Fatalf("error in benchmark setup: BenchmarkGarden_Plants requires valid garden")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, l := range test5.lookups {
			g.Plants(l.child)
		}
	}
}
