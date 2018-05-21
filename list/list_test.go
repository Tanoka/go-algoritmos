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

	if li.Val != 1 {
		t.Errorf("Error list current: value Expected:1, Actual:%d", li.Val)
	}

	if li.Length() != 3 {
		t.Errorf("Error length: Expected:3, Actual:%d", li.Length())
	}
}

func TestAddEnd(t *testing.T) {
	li := list.List{}
	li.AddEnd(1)
	li.AddEnd(2)
	li.AddEnd(3)

	if li.Val != 1 {
		t.Errorf("Error list current: value Expected:1, Actual:%d", li.Val)
	}

	if li.Length() != 3 {
		t.Errorf("Error length: Expected:3, Actual:%d", li.Length())
	}
}

func TestGoTo(t *testing.T) {
	li := &list.List{}
	li.AddBegin(1)
	li.AddBegin(2)
	li.AddBegin(3)

	li.GoBegin()
	if li.Val != 3 {
		t.Errorf("Error going to the list begin: value Expected:3, Actual:%d", li.Val)
	}

	li.GoEnd()
	if li.Val != 1 {
		t.Errorf("Error going to list end: value Expected:1, Actual:%d", li.Val)
	}
}

func TestNext(t *testing.T) {
	li := &list.List{}
	li.AddEnd(1)
	li.AddEnd(2)
	li.AddEnd(3)

	var err error
	sol := 1
	for ; err == nil; err = li.Next() {
		if li.Val != sol {
			t.Errorf("Error traversing list: value Expected:%d, Actual:%d", sol, li.Val)
		}
		sol++
	}

	if sol != 4 || li.Next() == nil {
		t.Errorf("Error traversing list, not end reached: Actual:%d, sol:%d", li.Val, sol)
	}
}

func TestNextEmpty(t *testing.T) {
	li := &list.List{}

	if err := li.Next(); err == nil {
		t.Errorf("Error traversing empty list")
	}
}

func TestPopEmptyList(t *testing.T) {
	li := &list.List{}

	if _, err := li.Pop(); err == nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}
}

func TestPopOneElement(t *testing.T) {
	li := &list.List{}
	li.AddEnd(7)

	value, err := li.Pop()

	if err != nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if value != 7 {
		t.Errorf("Error value pop. Expected:7 Actual:%d", value)
	}

	if li.Length() != 0 {
		t.Errorf("Error length after pop. Expected:0 Actual:%d", li.Length())
	}

	//Try Pop again. List should be empty
	_, err = li.Pop()
	if err.Error() != "Empty list" {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if li.Length() != 0 {
		t.Errorf("Error length after pop. Expected:0 Actual:%d", li.Length())
	}
}

func TestPopElement(t *testing.T) {
	li := &list.List{}
	li.AddBegin(7)
	li.AddBegin(11)
	li.AddBegin(13)
	li.AddBegin(17)

	value, err := li.Pop()

	if err != nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if value != 7 {
		t.Errorf("Error value pop. Expected:7 Actual:%d", value)
	}

	if li.Length() != 3 {
		t.Errorf("Error length after pop. Expected:3 Actual:%d", li.Length())
	}

	value, err = li.Pop()
	if err != nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if value != 11 {
		t.Errorf("Error value pop. Expected:11 Actual:%d", value)
	}

	if li.Length() != 2 {
		t.Errorf("Error length after pop. Expected:2 Actual:%d", li.Length())
	}
}

func TestShiftElement(t *testing.T) {
	li := &list.List{}
	li.AddBegin(7)
	li.AddBegin(11)
	li.AddBegin(13)
	li.AddBegin(17)

	value, err := li.Shift()

	if err != nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if value != 17 {
		t.Errorf("Error value pop. Expected:17 Actual:%d", value)
	}

	if li.Length() != 3 {
		t.Errorf("Error length after pop. Expected:3 Actual:%d", li.Length())
	}

	value, err = li.Shift()
	if err != nil {
		t.Errorf("Error pop empty list. Error:%v", err)
	}

	if value != 13 {
		t.Errorf("Error value pop. Expected:13 Actual:%d", value)
	}

	if li.Length() != 2 {
		t.Errorf("Error length after pop. Expected:2 Actual:%d", li.Length())
	}
}
