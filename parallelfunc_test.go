package parallelfunc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	jobs := make([]func() error, 0, 100)

	for i := 0; i < 20; i++ {

		a := func() error {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Printf("Sleeping %d seconds...\n", n)
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("Job\n")
			return nil
		}

		if i == 5 || i == 9 {
			jobs = append(jobs, func() error {
				return fmt.Errorf("Job error")
			})
		}

		jobs = append(jobs, a)
	}

	DoJobs(jobs, 5, 3)

}
