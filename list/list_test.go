package list

import "testing"

func TestGoBegin(t *testing.T) {
	li1 := List{1, nil, nil}
	li2 := List{2, nil, nil}
	li3 := List{3, nil, nil}

	li1.next = &li2
	li2.prev = &li1
	li2.next = &li3
	li3.prev = &li2

	bgn := GoBegin(&li3)

	if bgn.val != 1 {
		t.Errorf("Expected:1, Actual:%d", bgn.val)
	}
}

func TestLength(t *testing.T) {
	li1 := List{1, nil, nil}
	li2 := List{2, nil, nil}
	li3 := List{3, nil, nil}
	li4 := List{4, nil, nil}

	li1.next = &li2
	li2.prev = &li1
	li2.next = &li3
	li3.prev = &li2
	li3.next = &li4
	li4.prev = &li3

	cnt := Length(&li2)

	if cnt != 4 {
		t.Errorf("Expected:5, Actual:%d", cnt)
	}
}

func TestAddBegin(t *testing.T) {
	li := &List{1, nil, nil}

	AddBegin(li, 2)
	AddBegin(li, 3)
	AddBegin(li, 4)

	li = GoBegin(li)

	var x = 4
	for li.next != nil {
		if li.val != x {
			t.Errorf("Expected:%d, Actual:%d", x, li.val)
		}
		li = li.next
		x--
	}
}
