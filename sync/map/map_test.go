package main

import (
	"sync"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	rwLock  := sync.RWMutex{}
	mapA := make(map[int]int)

	for i:= 0; i < b.N; i++ {
		// rwLock.Lock()
		// mapA[i] = i
		// rwLock.Unlock()
		rwLock.RLock()
		_ = mapA[i]
		rwLock.RUnlock()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	mapB  := sync.Map{}
	for i := 0; i < b.N; i++ {
		//mapB.Store(i,i)
		_, ok := mapB.Load(i)
		if ok {
			
		}
	}
}