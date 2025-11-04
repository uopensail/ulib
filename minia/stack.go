package minia

/**
 * @brief Thread-safe generic stack implementation
 * @tparam T Element type
 */
type Stack[T any] struct {
	array []T
}

/**
 * @brief Create new stack instance
 * @tparam T Element type
 * @return New stack instance
 */
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array: make([]T, 0),
	}
}

/**
 * @brief Push element onto stack
 * @param val Element to push
 */
func (s *Stack[T]) Push(val T) {
	s.array = append(s.array, val)
}

/**
 * @brief Pop element from stack
 * @return Popped element
 * @throws panic if stack is empty
 */
func (s *Stack[T]) Pop() T {

	if len(s.array) == 0 {
		panic("stack is empty")
	}

	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val
}

/**
 * @brief Get stack size
 * @return Number of elements in stack
 */
func (s *Stack[T]) Size() int {
	return len(s.array)
}

/**
 * @brief Check if stack is empty
 * @return True if stack is empty
 */
func (s *Stack[T]) Empty() bool {
	return len(s.array) == 0
}

/**
 * @brief Top at top element without removing it
 * @return Top element
 * @throws panic if stack is empty
 */
func (s *Stack[T]) Top() T {
	if len(s.array) == 0 {
		panic("stack is empty")
	}

	return s.array[len(s.array)-1]
}

/**
 * @brief Clear all elements from stack
 */
func (s *Stack[T]) Clear() {
	s.array = s.array[:0]
}
