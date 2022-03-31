package set_test

import (
	"testing"

	"github.com/PacificBlackDuck/set"
)

func Test123(t *testing.T) {
	set := set.New[int]()
	set.Add(2)
	set.Add(2)
	set.Add(4)
	set.Add(6)

	if set.Len() != 3 {
		t.Fail()
	}

	set.Remove(2)
	set.Remove(2)
	set.Remove(2)

	if set.Len() != 2 {
		t.Fail()
	}
}

func TestLoop(t *testing.T) {
	set := set.NewFromSlice([]string{"hello", "world", "hello", "world", "testing", "123"})
	var i int
	for i = range set.Items() {
	}
	if i != 3 {
		t.Fail()
	}
}

func TestDeleting(t *testing.T) {
	set := set.NewFromSlice([]string{"hello", "world", "hello", "world", "testing", "123"})
	if set.Len() != 4 {
		t.Fail()
	}
	for _, item := range set.Items() {
		if item == "testing" {
			set.Remove(item)
		}
	}
	if set.Len() != 3 {
		t.Fail()
	}
}
