package list

type List struct {
	val  int
	next *List
	prev *List
}

func GoBegin(l *List) *List {
	for l.prev != nil {
		l = l.prev
	}
	return l
}

func AddBegin(l *List, x int) {
	node := &List{x, nil, nil}
	l = GoBegin(l)
	l.prev = node
	node.next = l
}

func Length(l *List) uint {
	l = GoBegin(l)

	var x uint = 1
	for ; l.next != nil; x++ {
		l = l.next
	}
	return x

}
