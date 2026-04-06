package list

type Node[K comparable, V any] struct {
	Key        K
	Value      V
	Prev, Next *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key:   key,
		Value: value,
	}
}

type LinkedList[K comparable, V any] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
}

func New[K comparable, V any]() *LinkedList[K, V] {
	return &LinkedList[K, V]{}
}

func (l *LinkedList[K, V]) prepend(node *Node[K, V]) {
	if l.Empty() {
		l.Head, l.Tail = node, node
		return
	}
	l.Head.Prev = node
	node.Next = l.Head
	l.Head = node
}

func (l *LinkedList[K, V]) append(node *Node[K, V]) {
	if l.Empty() {
		l.Head, l.Tail = node, node
		return
	}
	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
}

func (l *LinkedList[K, V]) PushFront(key K, value V) {
	node := NewNode(key, value)
	l.prepend(node)
}

func (l *LinkedList[K, V]) PushBack(key K, value V) {
	node := NewNode(key, value)
	l.append(node)
}

func (l *LinkedList[K, V]) PopFront() (K, V, bool) {
	if l.Empty() {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	first := l.Head
	l.Head = first.Next
	l.Head.Prev = nil
	return first.Key, first.Value, true
}

func (l *LinkedList[K, V]) PopBack() (K, V, bool) {
	if l.Empty() {
		var zeroKey K
		var zeroValue V
		return zeroKey, zeroValue, false
	}

	last := l.Tail
	l.Tail = last.Prev
	l.Tail.Next = nil
	return last.Key, last.Value, true
}

func (l *LinkedList[K, V]) MoveToFront(node *Node[K, V]) {
	l.Remove(node)
	l.prepend(node)
}

func (l *LinkedList[K, V]) MoveToBack(node *Node[K, V]) {
	l.Remove(node)
	l.append(node)
}

func (l *LinkedList[K, V]) Remove(node *Node[K, V]) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}
}

func (l *LinkedList[K, V]) Empty() bool {
	return l.Head == nil
}
