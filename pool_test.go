package main

import (
	"strings"
	"testing"
)

func TestNewPool(t *testing.T) {
	if p := newPool(); p == nil {
		t.Error("newPool(\"\") failed, got nil")
	}
}

func TestHandle(t *testing.T) {
	p := newPool()

	jobs := make(chan job)
	results := make(chan string)

	go p.handle(jobs, results)

	job, _ := newJob("invalid")
	jobs <- *job
	defer close(jobs)

	result := <-results
	if !strings.HasSuffix(result, "no such host") {
		t.Error("handle failed, expected no such host as suffix")
	}
}
