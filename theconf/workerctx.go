package main // OMIT

import ( // OMIT
	"context" // OMIT
	"sync"    // OMIT
) // OMIT

func worker(ctx context.Context, mainWG *sync.WaitGroup,
	input <-chan int, output chan<- int,
	parallelism int) {

	defer mainWG.Done()    // OMIT
	wg := sync.WaitGroup{} // OMIT
	wg.Add(parallelism)    // OMIT
	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg.Done()
			for { // HL
				select { // HL
				case <-ctx.Done(): // HL
					return
				case i := <-input:
					// ...
					output <- 1
				}
			}
		}()
	}
	wg.Wait()     // OMIT
	close(output) // OMIT
}
