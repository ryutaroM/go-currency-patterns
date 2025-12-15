package main

import "sync"

type ReadWriteMutex struct {
	readersCounter int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()
	for rw.writersWaiting > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.readersCounter++
	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriterLock() {
	rw.cond.L.Lock()

	rw.writersWaiting++
	for rw.readersCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.writersWaiting--
	rw.writerActive = true

	rw.cond.L.Unlock()
}

func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}
