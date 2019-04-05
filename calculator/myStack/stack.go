package myStack

type Stack struct {
	top []interface{}
	len int
}
func New()(*Stack){
	s := &Stack{}
	s.len = -1
	s.top = make([]interface{},20)
	return s
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.len
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top[s.len] = value
	s.len++
}

// Remove the top element from the stack
func (s *Stack) Pop() (value interface{}) {
	value = s.top[s.len]
	s.len--
	return value
}

//check if the stack is empty
func (s *Stack) IsEmpty() bool{
	return s.len == 0
}

func (s *Stack) Top() (value interface{}){
	if s.IsEmpty(){
		return nil
	}
	value = s.top[s.len]
	return 
}
