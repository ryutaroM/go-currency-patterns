package main

import "sync"

type ReaderWriterMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReaderWriterMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock() //block writers
	}
	rw.readersLock.Unlock()
}

func (rw *ReaderWriterMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock() //unblock writers
	}
	rw.readersLock.Unlock()
}

func (rw *ReaderWriterMutex) WriteLock() {
	rw.globalLock.Lock() // block readers and writers
}

func (rw *ReaderWriterMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReaderWriterMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}
