package main

import "sync"

func coordinate() {
	chan1 := make(chan interface{})
	chan2 := make(chan int)
	chan3 := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go worker1(&wg, chan1, chan2)
	wg.Add(1)
	go worker2(&wg, chan2, chan3)
	wg.Add(1)
	go worker3(&wg, chan3)

	wg.Wait()
}
