package xsync_test

import (
	"bytes"
	"testing"

	"github.com/playgroundgo/genlib/xsync"
)

func TestPool(t *testing.T) {
	bufPool := xsync.NewPool(func() *bytes.Buffer {
		return &bytes.Buffer{}
	})

	buf1 := bufPool.Get()
	bufPool.Put(buf1)
	buf2 := bufPool.Get()
	if buf1 != buf2 {
		t.Fatal("expected to get the same object from the pool")
	}

	buf3 := bufPool.Get()
	if buf3 == buf1 || buf3 == buf2 {
		t.Fatal("expected buf3 to be unique")
	}
}
