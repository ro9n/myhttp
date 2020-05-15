package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	const (
		HTTP  = "http://adjust.com/"
		HTTPS = "https://adjust.com/"
		EMPTY = "adjust.com/"
	)

	if _, err := newJob(""); err == nil {
		t.Error("newJob(\"\") failed, got nil")
	}

	if job, _ := newJob(EMPTY); job.url != HTTP {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", EMPTY, HTTP, job.url)
	}

	if job, _ := newJob(HTTP); job.url != HTTP {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", HTTP, HTTP, job.url)
	}

	if job, _ := newJob(HTTPS); job.url != HTTPS {
		t.Errorf("newJob(\"%s\") failed, expected %v, got %v", HTTPS, HTTPS, job.url)
	}
}

func TestExec(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}

	slow := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(5 * time.Second)
		_, _ = rw.Write([]byte(`OK`))
	}))

	defer slow.Close()

	j, _ := newJob(slow.URL)
	if _, err := j.exec(client); err == nil {
		t.Errorf("exec(\"%s\") failed, got nil", slow.URL)
	}

	fast := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(`OK`))
	}))

	defer fast.Close()

	j, _ = newJob(fast.URL)

	result, _ := j.exec(client)

	if len(result) == 0 {
		t.Errorf("exec(\"%s\") failed, expected non empty string, got empty", fast.URL)
	}

	words := strings.Split(result, " ")

	if len(words) != 2 {
		t.Errorf("exec(\"%s\") failed, expected  2 words, got %v", fast.URL, words)
	}

	expected := "e0aa021e21dddbd6d8cecec71e9cf564"
	if words[1] != expected {
		t.Errorf("exec(\"%s\") failed, expected  %v, got %v", fast.URL, expected, result)
	}

}
