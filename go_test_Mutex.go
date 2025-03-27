package main

import "sync"

//并发读写

type SyncMap struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func main() {

}
