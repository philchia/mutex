package mutex

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	state int32
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, 0, 1)
}

func (m *Mutex) Lock() {
	for !m.TryLock() {
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}
