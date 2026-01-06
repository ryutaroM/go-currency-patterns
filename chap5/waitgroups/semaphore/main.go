package main

import (
	listing5_16 "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter5/listing5.16"
)

type WaitGrp struct {
	sema *listing5_16.Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{sema: listing5_16.NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	wg.sema.Release()
}
