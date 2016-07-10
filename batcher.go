package batcher

import (
	"sync"
	"time"
)

type doFunc func(batch []interface{})

// Batcher is a simple batch utility to batch worky by a time duration.
type Batcher struct {
	timeout time.Duration
	lock    sync.Mutex
	do      doFunc
	input   chan interface{}
}

// New constructs a Batcher with the given duration and function.
func New(timeout time.Duration, do doFunc) *Batcher {
	b := &Batcher{
		timeout: timeout,
		do:      do,
	}

	return b
}

// Batch batches the given item that will be worked on.
func (b *Batcher) Batch(x interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.input == nil {
		b.input = make(chan interface{}, 4)

		go b.batch()
	}

	b.input <- x
}

func (b *Batcher) batch() {
	var batch []interface{}

	go b.sleeper()

	for item := range b.input {
		batch = append(batch, item)
	}

	b.do(batch)
}

func (b *Batcher) sleeper() {
	time.Sleep(b.timeout)

	b.lock.Lock()
	defer b.lock.Unlock()

	close(b.input)
	b.input = nil
}
