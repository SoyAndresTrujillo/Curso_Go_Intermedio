package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Job represents the job to be run
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Worker A buffered channel that we can send work requests on
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// Dispatcher NewWorker creates a new worker with the given id
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				// do the job
				fmt.Printf("Worker with id %d started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d finished with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				// we have been asked to stop
				fmt.Printf("Worker with id %d stopped\n", w.Id)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		WorkerPool: workerPool,
		MaxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQeue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// Parse the delay
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the value
	value, err := strconv.Atoi(r.FormValue("value")) // string to int
	if err != nil {
		http.Error(w, "Bad value: "+err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	jobQeue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	// create the job queue
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
