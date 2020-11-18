package monitormutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const rwmutexMaxReaders = 1 << 30

type RWMutex struct {
	sync.RWMutex
}

func NewRWMutex() *RWMutex {
	return &RWMutex{}
}

func (m *RWMutex) Count() int32 {
	readerCount := atomic.LoadInt32((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m.RWMutex)) + unsafe.Sizeof(sync.Mutex{}) + 2*unsafe.Sizeof(uint32(0)))))

	if readerCount >= 0 {
		return readerCount
	}

	return readerCount + rwmutexMaxReaders + 1 // +1 还有写锁 如果有多个写锁，会先调用写锁之间的互斥锁
}