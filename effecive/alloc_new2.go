package main

import (
	"fmt"
	"sync"
	"bytes"
)

type SyncedBuffer struct {
	Lock sync.Mutex
	Buffer bytes.Buffer
}

func main() {
	p := new(SyncedBuffer)
	fmt.Printf("%+v\n", p)
	var p2 SyncedBuffer
	fmt.Printf("%+v\n", p2)
}
