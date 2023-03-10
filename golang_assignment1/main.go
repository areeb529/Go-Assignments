package main

import (
	"fmt"
	"sync"
)

type video struct {
	id    string
	count int
	m     sync.Mutex
}

func (v *video) increment(wg *sync.WaitGroup) {
	v.m.Lock()
	v.count = v.count + 1
	v.m.Unlock()
	defer wg.Done()
}

func (v *video) viewCount(wg *sync.WaitGroup) {
	v.m.Lock()
	fmt.Printf("Video with id %s has %d view counts\n", v.id, v.count)
	v.m.Unlock()
	defer wg.Done()
}

func main() {
	var v1 = video{id: "gs62b3ys2", count: 0}
	
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go v1.viewCount(&wg)
		wg.Add(1)
		go v1.increment(&wg)
	}
	wg.Wait()
	

}
