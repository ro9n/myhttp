package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	concurrency := flag.Int("parallel", 10, "specify number of concurrent requests to send. "+
		"defaults to 10")

	flag.Parse()
	payload := flag.Args()

	if *concurrency <= 0 {
		log.Fatal("invalid argument: concurrency")
	}

	pool := newPool()
	jobs := make(chan job, len(payload))
	results := make(chan string, len(payload))

	for i := 0; i < *concurrency; i++ {
		go pool.handle(jobs, results)
	}

	defer close(jobs)

	for _, j := range payload {
		job, _ := newJob(j)
		jobs <- *job
	}

	defer close(jobs)

	for range payload {
		fmt.Println(<-results)
	}

	os.Exit(0)
}
