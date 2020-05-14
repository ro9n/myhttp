package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type job struct {
	url string
}

func newJob(url string) *job {
	j := new(job)
	j.url = compose(url)

	return j
}

func (j *job) exec(client *http.Client) (string, error) {
	resp, err := client.Get(j.url)
	if err != nil {
		return "", err
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	digest, err := checksum(string(text))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", j.url, digest), nil
}
