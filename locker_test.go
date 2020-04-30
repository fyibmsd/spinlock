package spinlock

import (
	"sync"
	"testing"
)

func TestSpinlock(t *testing.T) {
	l := Locker{}

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int, l *Locker) {
			l.Lock()
			defer l.Unlock()
			t.Logf("%d get lock", id)
			wg.Done()
		}(i, &l)
	}

	wg.Wait()
}
