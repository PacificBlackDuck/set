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

func Test456(t *testing.T) {
	set := set.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50)
	if set.Len() != 50 {
		t.Fail()
	}
	if set.Has(0) {
		t.Fail()
	}
	if !set.Has(33) {
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
