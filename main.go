package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	concurrency := flag.Int("parallel", 10, "specify number of concurrent requests to send. "+
		"defaults to 10")

	flag.Parse()
	payload := flag.Args()

	pool := newPool()
	jobs := make(chan job, len(payload))
	results := make(chan string, len(payload))

	for i := 0; i < *concurrency; i++ {
		go pool.handle(jobs, results)
	}

	for _, j := range payload {
		jobs <- *newJob(j)
	}

	close(jobs)

	for range payload {
		fmt.Println(<-results)
	}

	os.Exit(0)
}
