package monitormutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

// 监控mutex上等待的goroutine的数量
type Mutex struct {
	sync.Mutex
}

func NewMutex() *Mutex {
	return &Mutex{}
}

func (m *Mutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift //得到等待者的数值
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}