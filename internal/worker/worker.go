package worker

type Worker interface {
	Do(exec func(j int) (int, error)) error
}
