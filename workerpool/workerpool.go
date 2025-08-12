package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	numJobs    = 100
	numWorkers = 5
)

func main() {

	jobQueue := make(chan int)
	results := make(chan int)

	go startWorkerPool(jobQueue, results)
	time.Sleep(1 * time.Second)
	go beginJobStream(jobQueue)
	readResults(results)
}

func startWorkerPool(jobQueue <-chan int, results chan<- int) []*worker {
	fmt.Printf("\nStarting worker pool.\n")
	var workers []*worker
	for i := range numWorkers {
		speed := int(math.Max(float64(1), float64(rand.Intn(5))))
		w := newWorker(uint(i), speed, jobQueue, results, processFn)
		go w.start()
		workers = append(workers, w)
	}
	return workers
}

type worker struct {
	id             uint
	tasksCompleted uint
	speed          int
	inputStream    <-chan int
	outputStream   chan<- int
	processFunc    func(int, int) int
}

func newWorker(id uint, speed int, inputStream <-chan int, outputStream chan<- int, processFunc func(int, int) int) *worker {
	return &worker{
		id:             id,
		tasksCompleted: 0,
		speed:          speed,
		inputStream:    inputStream,
		outputStream:   outputStream,
		processFunc:    processFunc,
	}
}

func (w *worker) start() {
	fmt.Printf("\nStarting worker %v with speed %v.\n", w.id, w.speed)
	for input := range w.inputStream {
		output := w.processFunc(input, w.speed)
		w.outputStream <- output
		fmt.Printf("Worker %v processed an item: %v.\n", w.id, input)
	}
}

func processFn(input, speed int) int {
	time.Sleep(time.Second / time.Duration(speed*10))
	return input * 2
}

func beginJobStream(jobQueue chan<- int) {
	fmt.Printf("\nStarting job stream.\n\n")
	for i := range numJobs {
		jobQueue <- i
	}
}

func readResults(results <-chan int) {
	for i := range numJobs {
		result := <-results
		fmt.Printf("Result %v: %v\n", i, result)
	}
	fmt.Printf("\nAll jobs complete.\n")
}
