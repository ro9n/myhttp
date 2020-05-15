# myhttp
myhttp is a simple tool that generates  md5 digest of the request urls.
## Usage
#### Build
Build the project by running the following command. 
```sh
go build -o myhttp
```
#### Run
Set the *concurrency* parameter of the tool using the `parallel` flag
```sh
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```
Running without the parallel flag will result in 10 concurrent requests.
```sh
./myhttp http://www.adjust.com http://google.com
```
#### Test
To run all the unit tests, use the following command
```sh
go test
```
To show results for each test case set the verbose flag
```sh
go test -v
```
## Components
- `main.go` Driver component, represents the top layer 
- `pool.go` Represents a worker pool, abstraction for handling jobs concurrently.
- `job.go`  Abstraction for issuing a command to the worker pool, encapsulates execution logic.
- `hash.go` Responsible for generating md5 digest of a text.
- `url.go` Helper for appending protocol information to url.  
