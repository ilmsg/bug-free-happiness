//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// allocation with new and make
var v SyncedBuffer // type SyncedBuffer

func main() {
	p := new(SyncedBuffer) // type *SyncedBuffer
	fmt.Printf("%+v\n", p)

	v = SyncedBuffer{
		lock:   sync.Mutex{},
		buffer: bytes.Buffer{},
	}
	fmt.Printf("%+v\n", &v)
}
