package list_test

import "testing"
import "github.com/tanoka/go-algoritmos/list"

func TestLengthEmpty(t *testing.T) {
	li := list.List{}

	if li.Length() != 0 {
		t.Errorf("Expected:0, Actual:%d", li.Length())
	}
}

func TestAddBegin(t *testing.T) {
	li := list.List{}

	li.AddBegin(1)
	li.AddBegin(2)
	li.AddBegin(3)

        if li.Length() != 3 {
                t.Errorf("Expected:3, Actual:%d", li.Length())
        }
}

