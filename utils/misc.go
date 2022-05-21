package utils

import (
	"sync/atomic"
	"time"
)

type Reference struct {
	CloseHandler   func()
	referenceCount int32
}

func (ref *Reference) Retain() {
	atomic.AddInt32(&ref.referenceCount, 1)
}
func (ref *Reference) Release() {
	atomic.AddInt32(&ref.referenceCount, -1)
}

func (ref *Reference) Free() {
	for atomic.LoadInt32(&ref.referenceCount) > 0 {
		time.Sleep(time.Second)
	}
	if ref.CloseHandler != nil {
		ref.CloseHandler()
	}
}
