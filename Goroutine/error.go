package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

type TaskResult struct {
	TaskID int
	Output string
	Error  error
}

type Pool struct {
	workers  int
	taskCh   chan Task
	resultCh chan TaskResult
	errCh    chan error // For pool-level errors
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewPool(workers int, bufferSize int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		workers:  workers,
		taskCh:   make(chan Task, bufferSize),
		resultCh: make(chan TaskResult, bufferSize),
		errCh:    make(chan error, 1),
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (p *Pool) Start() {
	p.wg.Add(p.workers)
	for i := 0; i < p.workers; i++ {
		go p.worker(i)
	}

	// Close results when all workers done
	go func() {
		p.wg.Wait()
		close(p.resultCh)
	}()
}

func (p *Pool) worker(id int) {
	defer p.wg.Done()

	for {
		select {
		case task, ok := <-p.taskCh:
			if !ok {
				log.Printf("Worker %d: tasks channel closed", id)
				return
			}

			// Simulate work with potential error
			output, err := p.process(task)

			// Try to send result with timeout
			select {
			case p.resultCh <- TaskResult{
				TaskID: task.ID,
				Output: output,
				Error:  err,
			}:
			case <-p.ctx.Done():
				log.Printf("Worker %d: context cancelled", id)
				return
			case <-time.After(5 * time.Second):
				log.Printf("Worker %d: timeout sending result", id)
				p.errCh <- fmt.Errorf("worker %d: send timeout", id)
				return
			}

		case <-p.ctx.Done():
			log.Printf("Worker %d: context cancelled", id)
			return
		}
	}
}

func (p *Pool) process(t Task) (string, error) {
	if t.Data == "error" {
		return "", errors.New("simulated error")
	}
	if t.Data == "slow" {
		time.Sleep(2 * time.Second)
	}
	return fmt.Sprintf("processed-%s", t.Data), nil
}

func (p *Pool) Submit(t Task) error {
	select {
	case p.taskCh <- t:
		return nil
	case <-p.ctx.Done():
		return errors.New("pool closed")
	case <-time.After(time.Second):
		return errors.New("submit timeout")
	}
}

func (p *Pool) Results() <-chan TaskResult {
	return p.resultCh
}

func (p *Pool) Stop() {
	p.cancel()
	close(p.taskCh)
}

func main() {
	pool := NewPool(3, 10)
	pool.Start()

	// Submit tasks
	go func() {
		for i := 0; i < 20; i++ {
			data := fmt.Sprintf("task-%d", i)
			if i == 5 {
				data = "error" // Trigger error
			}

			if err := pool.Submit(Task{ID: i, Data: data}); err != nil {
				log.Printf("Submit failed: %v", err)
				return
			}
		}
	}()

	// Collect results
	var success, failed int
	for res := range pool.Results() {
		if res.Error != nil {
			log.Printf("Task %d failed: %v", res.TaskID, res.Error)
			failed++
		} else {
			fmt.Printf("Task %d success: %s\n", res.TaskID, res.Output)
			success++
		}
	}

	fmt.Printf("\nDone: %d success, %d failed\n", success, failed)
	pool.Stop()
}
