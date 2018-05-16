package list

type cell struct {
	val  int
	next *cell
	prev *cell
}

type List struct {
	*cell
}

func (l *List)GoBegin() *List {
	if l.cell == nil  {
		return l
	}
	for l.cell.prev != nil {
		l.cell = l.cell.prev
	}
	return l
}

func (l *List)AddEnd(x int) {
	node := &cell{x, nil, nil}
	for l.cell.next != nil {
		l.cell = l.cell.next
	}
	l.cell.next = node
	node.prev = l.cell
}

func (l *List)AddBegin(x int) {
	node := &cell{x, nil, nil}
	l = l.GoBegin()
	if l.cell != nil {
		l.cell.prev = node
		node.next = l.cell
	} else {
		l.cell = node
	}
}

func (l *List)Length() uint {
	if l.cell == nil {
		return 0
	}
	l = l.GoBegin()
	var x uint = 1
	for ; l.cell.next != nil; x++ {
		l.cell = l.cell.next
	}
	return x

}
