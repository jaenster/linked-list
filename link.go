package list

type Link[T any] struct {
	value *T
	next  *Link[T]
	prev  *Link[T]
}

func New[T any](v *T) *Link[T] {
	return &Link[T]{
		value: v,
		next:  nil,
		prev:  nil,
	}
}
func Empty[T any]() *Link[T] {
	return &Link[T]{
		value: nil,
		next:  nil,
		prev:  nil,
	}
}

func (l *Link[T]) Append(v *T) *Link[T] {
	if l.value == nil {
		l.value = v
	} else if l.next == nil {
		l.next = &Link[T]{value: v, next: nil, prev: l}
	} else {
		l.next.Append(v)
	}
	return l
}

func (l *Link[T]) PushBack(v *T) *Link[T] {
	return l.Append(v)
}

func (l *Link[T]) Prepend(v *T) *Link[T] {
	if l.value == nil {
		l.value = v
	} else if l.prev == nil {
		l.prev = &Link[T]{value: v, next: l, prev: nil}
	} else {
		l.prev.Prepend(v)
	}
	return l
}

func (l *Link[T]) PushFront(v *T) *Link[T] {
	return l.Prepend(v)
}

func (l *Link[T]) Remove(v *T) *Link[T] {
	if l.next != nil {
		if l.next.value == v {
			l.next = l.next.next
			l.prev = l.next.prev
		} else {
			l.next.Remove(v)
		}
	}
	return l
}

func (l *Link[T]) Contains(v *T) bool {
	if l.value == v {
		return true
	} else if l.next != nil {
		if l.next.value == v {
			return true
		} else {
			return l.next.Contains(v)
		}
	}
	if l.prev != nil {
		if l.prev.value == v {
			return true
		} else {
			return l.prev.Contains(v)
		}
	}
	return false
}

func (l *Link[T]) Len() int {
	if l.next == nil {
		return 0
	} else {
		if l.value == nil {
			return 0 + l.next.Len()
		}
		return 1 + l.next.Len()
	}
}

func (l *Link[T]) Destroy() *Link[T] {
	if l.next != nil {
		l.next.Destroy()
	}
	if l.prev != nil {
		l.prev.Destroy()
	}
	l.value = nil
	l.next = nil
	l.prev = nil
	return l
}

func (l *Link[T]) Head() *Link[T] {
	if l.prev == nil {
		return l
	} else {
		return l.prev.Head()
	}
}

func (l *Link[T]) Tail() *Link[T] {
	if l.next == nil {
		return l
	} else {
		return l.next.Tail()
	}
}

func (l *Link[T]) GetNext() *Link[T] {
	return l.next
}

func (l *Link[T]) GetPrev() *Link[T] {
	return l.prev
}

func (l *Link[T]) GetValue() *T {
	return l.value
}

func (l *Link[T]) SetValue(v *T) {
	l.value = v
}

func (l *Link[T]) SetNext(n *Link[T]) {
	l.next = n
}

func (l *Link[T]) SetPrev(p *Link[T]) {
	l.prev = p
}

func (l *Link[T]) Loop(f func(v *T)) {
	if l.value != nil {
		f(l.value)
	}
	if l.next != nil {
		l.next.Loop(f)
	}
}

func (l *Link[T]) ReverseLoop(f func(v *T)) {
	if l.value != nil {
		f(l.value)
	}
	if l.prev != nil {
		l.prev.ReverseLoop(f)
	}
}

func (l *Link[T]) Emplace(v *T) *Link[T] {
	if l.value == nil {
		l.value = v
	} else if l.next == nil {
		l.next = &Link[T]{value: v, next: nil, prev: l}
	} else {
		l.next.Emplace(v)
	}
	return l
}
