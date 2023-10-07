package _3_concurrency

import (
	"sync"
	"testing"
)

func benchLock(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Read()
			}()
		}

		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Write()
			}()
		}

		wg.Wait()
	}

}

func BenchmarkLockReadMore(b *testing.B)   { benchLock(b, &Lock{}, 9, 1) }
func BenchmarkLockReadMoreRW(b *testing.B) { benchLock(b, &RWLock{}, 9, 1) }

func BenchmarkLockWriteMore(b *testing.B)   { benchLock(b, &Lock{}, 1, 9) }
func BenchmarkLockWriteMoreRW(b *testing.B) { benchLock(b, &RWLock{}, 1, 9) }

func BenchmarkLockEqual(b *testing.B)   { benchLock(b, &Lock{}, 5, 5) }
func BenchmarkLockEqualRW(b *testing.B) { benchLock(b, &RWLock{}, 5, 5) }
