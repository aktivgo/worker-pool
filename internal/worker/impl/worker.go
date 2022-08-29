package impl

import (
	"fmt"
)

type worker struct {
	jobs    <-chan int
	results chan<- int
}

func NewWorker(
	jobs <-chan int,
	results chan<- int,
) *worker {
	return &worker{
		jobs:    jobs,
		results: results,
	}
}

func (w *worker) Do(id int, process func(j int) (int, error)) error {
	for job := range w.jobs {
		fmt.Println("worker", id, "started  job", job)

		result, err := process(job)
		if err != nil {
			return err
		}

		fmt.Println("worker", id, "finished job", job)

		w.results <- result
	}

	return nil
}
