package app

import (
	"log"
	"time"
	"worker-pool/internal/worker/impl"
)

func Run() {
	const jobsCount = 16
	const workersCount = 16

	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	worker := impl.NewWorker(
		jobs,
		results,
	)

	for i := 0; i < workersCount; i++ {
		go worker.Do(i+1, processJob)
	}

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= jobsCount; i++ {
		log.Println(<-results)
	}
}

func processJob(j int) (int, error) {
	time.Sleep(time.Second)
	return j * 2, nil
}
