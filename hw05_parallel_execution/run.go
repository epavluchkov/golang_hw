package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tCh := make(chan Task)
	mu := sync.Mutex{}
	errCnt := 0
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tCh {
				err := t()
				if err != nil {
					mu.Lock()
					errCnt++
					mu.Unlock()
				}
			}
		}()
	}

	for _, t := range tasks {
		mu.Lock()
		if errCnt >= m && m != 0 {
			mu.Unlock()
			break
		}
		mu.Unlock()
		tCh <- t
	}
	close(tCh)

	wg.Wait()

	if errCnt >= m && m != 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
