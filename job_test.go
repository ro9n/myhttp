package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	const (
		http  = "http://adjust.com/"
		https = "https://adjust.com/"
		empty = "adjust.com/"
	)

	if _, err := newJob(""); err == nil {
		t.Error("newJob(\"\") failed, got nil")
	}

	if job, _ := newJob(empty); job.url != http {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", empty, http, job.url)
	}

	if job, _ := newJob(http); job.url != http {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", http, http, job.url)
	}

	if job, _ := newJob(https); job.url != https {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", https, https, job.url)
	}
}

func TestExec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(10 * time.Second)
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	j, _ := newJob(server.URL)
	if _, err := j.exec(client); err == nil {
		t.Error("exec(\"\") failed, got nil")
	}
}
