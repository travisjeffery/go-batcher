package batcher

import (
	"sync"
	"testing"
	"time"

	"sync/atomic"

	"github.com/bmizerany/assert"
)

func TestBatcherOk(t *testing.T) {
	var called int32

	b := New(5*time.Millisecond, func(batch []interface{}) error {
		atomic.AddInt32(&called, 1)
		assert.Equal(t, 10, len(batch))

		return nil
	})

	var wg sync.WaitGroup
	iters := 10

	wg.Add(iters)

	for i := 0; i < iters; i++ {
		go func(i int) {
			b.Batch(i)

			wg.Done()
		}(i)
	}

	wg.Wait()

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, int(called), 1)
}
