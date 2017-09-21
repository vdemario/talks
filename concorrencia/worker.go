package main // OMIT

import "sync"

func worker(mainWG *sync.WaitGroup, input <-chan int, output chan<- int, parallelism int) {

	defer mainWG.Done()

	wg := sync.WaitGroup{}
	wg.Add(parallelism)

	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg.Done()
			for i := range input {
				// ...
				output <- 1
			}
		}()
	}

	wg.Wait()
	close(output)
}
