package FibonacciHeap

type FibonacciHeap struct {
	rootList LinkedList
	min      *Node
}

func (fh *FibonacciHeap) GetMin() int {
	return fh.min.value
}

func (fh *FibonacciHeap) ExtractMin() int {
	value := fh.min.value

	for e := fh.min.childList.front; e != nil; e = e.next {
		fh.min.childList.remove(e)
		fh.rootList.insert(e)
		e.parent = nil
	}

	fh.rootList.remove(fh.min)
	if fh.rootList.front == nil {
		fh.min = nil
	} else {
		fh.cleanUp()
	}

	return value
}

func (fh *FibonacciHeap) Insert(val int) {
	newNode := new(Node)
	newNode.value = val
	newNode.rootList = &fh.rootList

	fh.insert(newNode)
}

func (fh *FibonacciHeap) insert(n *Node) {
	fh.rootList.insert(n)

	if fh.min == nil || fh.rootList.end.value < fh.min.value {
		fh.min = fh.rootList.end
	}
}

func (fh *FibonacciHeap) cleanUp() {
	m := make(map[int]*Node)
	rl := []*Node{}

	for e := fh.rootList.front; e != nil; e = e.next {
		rl = append(rl, e)
	}

	for _, n := range rl {
		n.prev = nil
		n.next = nil
		for true {
			degrees := n.GetDegrees()
			if m[degrees] != nil {
				n = m[degrees].combine(n)
				m[degrees] = nil
			} else {
				m[degrees] = n
				break
			}

		}
	}

	fh.rootList.front = nil
	fh.rootList.end = nil
	fh.min = nil

	for _, n := range m {
		if n != nil {
			fh.insert(n)
		}
	}
}

type LinkedList struct {
	front *Node
	end   *Node
}

func (l *LinkedList) insert(n *Node) {
	n.prev = l.end

	if l.front == nil {
		l.front = n
		l.end = n
	} else {
		if l.end != nil {
			l.end.next = n
		}
		l.end = n
	}

}

func (l *LinkedList) remove(n *Node) {
	if n == l.front {
		l.front = n.next
	}
	if n == l.end {
		l.end = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

type Node struct {
	value     int
	next      *Node
	prev      *Node
	parent    *Node
	childList LinkedList
	rootList  *LinkedList
	marked    bool
}

func (n *Node) GetDegrees() int {
	largestDegrees := 0
	for e := n.childList.front; e != nil; e = e.next {
		degrees := e.GetDegrees()
		if degrees > largestDegrees {
			largestDegrees = degrees
		}
	}

	return largestDegrees + 1
}

func (n *Node) SetValue(val int) {
	n.value = val
	if n.parent != nil && n.value < n.parent.value {
		n.cutOut()
	}
}

func (a *Node) combine(b *Node) *Node {
	parent := &Node{}
	child := &Node{}

	if a.value < b.value {
		parent = a
		child = b
	} else {
		parent = b
		child = a
	}

	a.rootList.remove(child)
	parent.childList.insert(child)
	child.parent = parent
	return parent
}

func (n *Node) cutOut() {
	if n.parent != nil {
		parent := n.parent

		n.parent.childList.remove(n)
		n.rootList.insert(n)
		n.parent = nil
		n.marked = false

		if parent.marked == true {
			parent.cutOut()
		} else {
			parent.marked = true
		}
	}
}
