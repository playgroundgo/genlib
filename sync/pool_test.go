package sync_test

import (
	"testing"

	"github.com/playgroundgo/genlib/sync"
)

func TestPool(t *testing.T) {
	bufPool := sync.NewPool(func() *[]byte {
		buf := make([]byte, 512*1024)
		return &buf
	})
	buf1 := bufPool.Get()
	bufPool.Put(buf1)
	buf2 := bufPool.Get()
	if buf1 != buf2 {
		t.Fatal("expected to get the same object from the pool")
	}
}
