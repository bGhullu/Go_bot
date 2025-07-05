package main

import (
	"fmt"
	"time"
)

type Task interface {
	Run()
}

type Job struct {
	ID int
}

func (j Job) Run() {
	fmt.Println("Running job:", j.ID)
}

func dispatch(tasks []Task) {
	for _, t := range tasks {
		go t.Run()
	}
}

func main() {
	jobs := []Task{
		Job{ID: 1},
		Job{ID: 2},
	}
	for _, job := range jobs {
		job.Run()
	}

	dispatch(jobs)
	time.Sleep(1 * time.Second)

}
