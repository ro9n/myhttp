package main

import (
	"fmt"
	"net/http"
	"time"
)

type pool struct {
	client *http.Client
}

func newPool() *pool {
	p := new(pool)
	p.client = &http.Client{Timeout: 5 * time.Second}

	return p
}

func (p *pool) handle(jobs <-chan job, results chan<- string) {
	for job := range jobs {
		result, err := job.exec(p.client)

		if err != nil {
			results <- fmt.Sprintf("%s %v", job.url, err.Error())
		} else {
			results <- result
		}
	}
}
