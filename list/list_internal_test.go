package list

import "testing"

func TestGoBegin(t *testing.T) {
	ce1 := cell{1, nil, nil}
	ce2 := cell{2, nil, nil}
	ce3 := cell{3, nil, nil}

	ce1.next = &ce2
	ce2.prev = &ce1
	ce2.next = &ce3
	ce3.prev = &ce2

	li := List{&ce3, 3, &ce1, &ce3}
	li.GoBegin()

	if li.cell.Val != 1 {
		t.Errorf("Expected:1, Actual:%d", li.cell.Val)
	}
}

func TestLengthEmpty(t *testing.T) {
	li := List{}

	if li.Length() != 0 {
		t.Errorf("Expected:0, Actual:%d", li.Length())
	}
}

func TestLength(t *testing.T) {
	ci1 := cell{1, nil, nil}
	ci2 := cell{2, nil, nil}
	ci3 := cell{3, nil, nil}
	ci4 := cell{4, nil, nil}

	ci1.next = &ci2
	ci2.prev = &ci1
	ci2.next = &ci3
	ci3.prev = &ci2
	ci3.next = &ci4
	ci4.prev = &ci3

	li := List{&ci2, 4, &ci1, &ci4}

	cnt := li.Length()

	if cnt != 4 {
		t.Errorf("Expected:5, Actual:%d", cnt)
	}
}
