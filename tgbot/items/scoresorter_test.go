package items

import (
	"testing"
)

var newScoreSorterTest = []struct {
	in     map[int]int64
	result []int
}{
	{nil, make([]int, 0)},
	{
		map[int]int64{
			1: 1,
			2: 2,
			3: 3,
		},
		[]int{
			1,
			2,
			3,
		},
	},
}

func TestNewScoreSorter(t *testing.T) {
	for _, nsst := range newScoreSorterTest {
		obj := NewScoreSorter(nsst.in)

		if obj.results == nil {
			t.Error("Found nil value for result fields, want []string")
			continue
		}

		if len(obj.userScores) != len(nsst.in) {
			t.Error("Length of user score mapper is not the same as the the given input mapper.")
			continue
		}

		if len(obj.results) != len(nsst.in) {
			t.Error("Length of result struct is not the same as the the given input mapper.")
			continue
		}

		for i, a := range nsst.in {
			if _, ok := obj.userScores[i]; !ok {
				t.Errorf("Couldn't find %d in user score object", i)
				continue
			}

			if a != obj.userScores[i] {
				t.Errorf("Got %d for user %d, want %d", obj.userScores[i], i, a)
			}
		}

		for _, b := range obj.results {
			if _, ok := nsst.in[b]; !ok {
				t.Errorf("Couldn't find %d in result struct", b)
				continue
			}
		}

	}
}

var lentest = []struct {
	in     map[int]int64
	length int
}{
	{nil, 0},
	{
		map[int]int64{
			1: 1,
			2: 2,
			3: 3,
		},
		3,
	},
}

func TestLen(t *testing.T) {
	for _, l := range lentest {
		obj := NewScoreSorter(l.in)

		if obj.Len() != l.length {
			t.Errorf("Got length %d from lenght method, want %d", obj.Len(), l.length)
		}
	}
}

// TODO: implement missing test function for Less, Swap and Sorted.
