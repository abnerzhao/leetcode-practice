package stack

import "sync"

// Item the type of stack
type Item interface{}

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// NewStack create a new ItemStack
func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []Items{}
	return s
}

// Push add Item to the top of stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop remove Item from the top of stack
func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Len length of stack
func (s ItemStack) Len() int {
	return len(s)
}

// IsEmpty whether the stack is empty
func (s ItemStack) IsEmpty() bool {
	return len(s) == 0
}

// GetTop get top of stack
func (s ItemStack) GetTop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}
