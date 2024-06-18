package MutexStruct

import "sync"

type MyStruct struct {
	mutex sync.Mutex
	value int
}

func (s *MyStruct) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.value++
}

func (s *MyStruct) IncrementBy(n int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.value += n
}

func (s *MyStruct) Value() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.value
}
