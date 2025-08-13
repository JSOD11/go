package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var (
	numJobs    = 100000
	numWorkers = 10
	speedLimit = 1000
	wg         sync.WaitGroup
)

func main() {

	jobQueue := make(chan int)
	results := make(chan int)

	workers := startWorkerPool(jobQueue, results)
	time.Sleep(time.Second / 2)

	go beginJobStream(jobQueue)

	wg.Add(1)
	go readResults(results)

	wg.Wait()

	summarizePerformance(workers)
}

func startWorkerPool(jobQueue <-chan int, results chan<- int) []*worker {
	fmt.Printf("Starting worker pool.\n")
	var workers []*worker
	for i := range numWorkers {
		speed := int(math.Max(float64(1), float64(rand.Intn(speedLimit))))
		w := newWorker(uint(i), speed, jobQueue, results, processFn)

		wg.Add(1)
		go w.start()

		workers = append(workers, w)
	}
	return workers
}

type worker struct {
	id                uint
	numTasksCompleted uint
	speed             int
	inputStream       <-chan int
	outputStream      chan<- int
	processFunc       func(int, int) int
}

func newWorker(
	id uint,
	speed int,
	inputStream <-chan int,
	outputStream chan<- int,
	processFunc func(int, int) int) *worker {
	return &worker{
		id:                id,
		numTasksCompleted: 0,
		speed:             speed,
		inputStream:       inputStream,
		outputStream:      outputStream,
		processFunc:       processFunc,
	}
}

func (w *worker) start() {
	defer wg.Done()
	for input := range w.inputStream {
		output := w.processFunc(input, w.speed)
		w.numTasksCompleted++
		w.outputStream <- output
	}
}

func processFn(input, speed int) int {
	time.Sleep(time.Second / time.Duration(speed*10))
	return input
}

func beginJobStream(jobQueue chan<- int) {
	fmt.Printf("Starting job stream.\n")
	for i := range numJobs {
		jobQueue <- i
	}
	close(jobQueue)
}

func readResults(results <-chan int) {
	defer wg.Done()
	for range numJobs {
		<-results
	}
}

func summarizePerformance(workers []*worker) {
	fmt.Printf("\n———————— Performance Summary ———————\n")
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].speed > workers[j].speed
	})
	for _, w := range workers {
		fmt.Printf("\nWorker %v\n——————————————————\nSpeed: %v\nNum tasks completed: %v\n", w.id, w.speed, w.numTasksCompleted)
	}
}
