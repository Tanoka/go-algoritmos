package list

import "errors"

type cell struct {
	Val        int
	next, prev *cell
}

type List struct {
	*cell
	length   uint
	ini, end *cell
}

func (l *List) GoBegin() {
	l.cell = l.ini
}

func (l *List) GoEnd() {
	l.cell = l.end
}

func (l *List) AddEnd(x int) {
	node := &cell{x, nil, nil}
	if l.ini == nil {
		l.cell = node
		l.ini = node
	} else {
		l.end.next = node
		node.prev = l.end
	}
	l.end = node
	l.length++
}

func (l *List) AddBegin(x int) {
	node := &cell{x, nil, nil}
	if l.ini == nil {
		l.cell = node
		l.end = node
	} else {
		l.ini.prev = node
		node.next = l.ini
	}
	l.ini = node
	l.length++
}

func (l *List) Length() uint {
	return l.length

}

func (l *List) Next() error {
	var err error
	if l.length == 0 {
		err = errors.New("Empty list")
	} else if l.cell.next != nil {
		l.cell = l.cell.next
	} else {
		err = errors.New("End list")
	}
	return err
}

func (l *List) Prev() error {
	var err error
	if l.length == 0 {
		err = errors.New("Empty List")
	} else if l.cell.prev != nil {
		l.cell = l.cell.prev
	} else {
		err = errors.New("Begin list")
	}
	return err
}

//Return last element of the list. If there is no elmenent return a error
func (l *List) Pop() (re int, err error) {
	if l.end != nil { //There is a last element
		re = l.end.Val
		//The prev element is now the last
		l.end = l.end.prev
		//Current cell is the last now
		l.cell = l.end
		l.length--
	} else {
		err = errors.New("Empty list")
	}
	return
}

//Return the first element of the list. If there is no elmenent return a error
func (l *List) Shift() (re int, err error) {
	if l.ini != nil {
		re = l.ini.Val
		l.ini = l.ini.next
		l.cell = l.ini
		l.length--
	} else {
		err = errors.New("Empty list")
	}
	return
}
