package parallelfunc

import (
	"fmt"
	"sync"
)

func DoJobs(jobs []func() error, maxRoutines, maxErrors int) {

	wg := &sync.WaitGroup{}
	locker := &sync.Mutex{}
	errorsMax := maxErrors
	errorsCount := 0
	jobsChan := make(chan func() error, len(jobs))
	stop := make(chan struct{})

	wg.Add(maxRoutines)
	for i := 0; i < maxRoutines; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-stop:
					return
				case job, ok := <-jobsChan:
					if !ok {
						return
					}
					if err := job(); err != nil {
						fmt.Println("Error :", err)
						locker.Lock()

						errorsCount++
						if errorsCount >= errorsMax {
							locker.Unlock()
							close(stop)
							return
						}
						locker.Unlock()
					}
				}
			}
		}()
	}

	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	wg.Wait()
}
