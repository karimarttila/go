# Go Simple Server  <!-- omit in toc -->


# Table of Contents  <!-- omit in toc -->
- [Introduction](#introduction)
- [Learning Process](#learning-process)
- [Go](#go)
- [GoLand](#goland)
- [Go Code Format](#go-code-format)
- [Static Code Analysis](#static-code-analysis)
- [Testing](#testing)
- [GoLand Debugger](#goland-debugger)
- [Test Performance Between Five Languages](#test-performance-between-five-languages)
- [Go REPL](#go-repl)
- [Logging](#logging)
- [Readability](#readability)
- [Productivity](#productivity)
- [Lines of Code](#lines-of-code)
- [Performance](#performance)
- [Conclusions](#conclusions)


# Introduction

It's really fun and interesting to learn a new programming language. When you have zero knowledge before taking the challenge learning the new language is a kind of exploration expedition to a new land with different language and customs - the only way to effectively learn it is to go there, learn the language and its idioms.

# Learning Process

I had practically zero knowledge of Go when I started this project. I watched one short "Go Basics" type Pluralsight video before starting to program Go. Mostly everthing I just learned on the fly while doing programming. E.g. I was wondering what to do in a situation in which some Go function returns three distinct return values but I need only one? I handled these situations just by googling, e.g.  ```go function returns multiple values "unused variable"``` => use underscore for those variables you don't need.

For my own learning purposes I commented quite a few of these lines so that I can use this Go code of Simple Server for my future reference implementation of Go. Example:

```go
// NOTE: In Go public variables and functions start with capital letter.
var MyLogger = initLogger()
...
// NOTE: Use underscore '_' when you don't need to reference certain return values.
pc, _, _, _ := runtime.Caller(1)
```

# Go

I was using [Go](https://golang.org//) 1.11 on Ubuntu18 when implementing this Simple Server.

You have to set the $GOPATH and $GOROOT environmental variables to point to your Go project directory and where your Go installation is. See example in (setenv.sh)[TODO].

```bash
go version      => go version go1.11.1 linux/amd64
pwd             => /mnt/edata/aw/kari/github/go
echo $GOPATH    => /mnt/edata/aw/kari/github/go
echo $GOROOT    => /mnt/local/go-1.11
```

I used [dep](https://github.com/golang/go/wiki/PackageManagementTools) tool to mangage Go packages:

```bash
dep init src/github.com/karimarttila/go/simpleserver
``` 

NOTE: This is  a one time task and creates the [Gopkg.toml](TODO) file.

Running ```dep ensure``` in the project (simpleserver) updates the [Gopkg.lock](TODO) file and creates a vendor directory (not in Git). 

I created a Git repo for this project and added it under the GOPATH in my machine. I also used the go get <url> to load the source code of the dependencies that I'm using (e.g. Logrus).

TODO: Read more about dep ensure and vendor!


And so we finalized our short tour to "Go and package management."


# GoLand

I use [GoLand](https://www.jetbrains.com/go/)

 PyCharm is really good IDE for Python programming - the editor is great and there are a lot of utilities that make your Python programming more productive (code completion, test runners, automatic linter ([PEP](https://www.python.org/dev/peps/pep-0008/)) etc). 

I use [IntelliJ IDEA](https://www.jetbrains.com/idea/) for Java programming, [PyCharm](https://www.jetbrains.com/pycharm) for Python programming and IntelliJ IDEA with [Cursive](https://cursive-ide.com/) plugin for Clojure programming. Since GoLand, PyCharm and IDEA are provided by the same company (JetBrains) they provide very similar look-and-feel. So, there are a lot of synergy benefits to use the same IDE for several programming languages.


# Go Code Format

Go is an interesting language in that sense that format of the Go code is very opinionated. Very opinionated in that sense that the Go compiler even requests code to be in certain format or it doesn't compile the code even though it would be syntacally right. Formatting of the Go code is build into the language (see: [format](https://golang.org/pkg/go/format/)).

You can reformat the Simple Server code using the following command in the $GOPATH directory.

```bash
go fmt github.com/karimarttila/go/simpleserver
```

I also provided a script [go-fmt-simpleserver.sh](https://github.com/karimarttila/go/tree/master/simpleserver/scripts) to run the fmt code formatter to all Simple Server Go files. It's a good idea to run this script every once in a while to keep the Go code formatting clean.


# Static Code Analysis

Go provides a simple static code analysis tool in the standard Go toolbox: [vet](https://golang.org/cmd/vet/). You can run the static code analysis using command: 

```bash
go vet github.com/karimarttila/go/simpleserver/...
```

I also found another interesting tool: [Staticcheck](https://staticcheck.io/docs/). Install the package and you can run staticcheck, gosimple and unused in one go for all Go code base files:

```bash
megacheck github.com/karimarttila/go/simpleserver/app/...
```

Staticcheck open source version is free. If you find the tool useful you should consider buying the commercial version. 


# Testing

Go is pretty amazing in that sense that you have also [the Go testing framework](https://golang.org/pkg/testing/) in the standard library.

Go tests are pretty easy to create. You don't have asserts but instead you just write standard application logic to test whether your package works as expected. Example:

```go
func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	myProductGroups := GetProductGroups()
	myPGMap := myProductGroups.ProductGroupsMap
	if len(myPGMap) != 2 {
		t.Errorf("There should be exactly two product groups, got: %d", len(myPGMap))
	}
	if myPGMap["1"] != "Books" || myPGMap["2"] != "Movies" {
		t.Errorf("There were wrong values for product groups: %s", myPGMap)
	}
}
```

You can run all tests in command line using command (TODO: update when all tests have been implemented):

```bash
go test github.com/karimarttila/go/simpleserver/...
ok  	github.com/karimarttila/go/simpleserver/app/domaindb	0.002s
ok  	github.com/karimarttila/go/simpleserver/app/main	0.003s
ok  	github.com/karimarttila/go/simpleserver/app/util	0.003s
ok  	github.com/karimarttila/go/simpleserver/app/webserver	0.002s
```

Running tests is pretty nice since Go compiles really fast and starts the tests immediately.

# GoLand Debugger

GoLand debugger is really good. Debugger starts immediately and is really fast. It's not a Lisp REPL but a pretty good second option. Go's data structures are pretty simple and GoLand debugger does a very good job presenting the data structures and values in the editor and in the variables window.


# Test Performance Between Five Languages

All right! Finally my journey travelling through five language lands is over and I can compare the languages. Let's compare the test performance between the languages.


**Clojure**:

```bash
time ./run-tests.sh 
19:52:18.637 [main] INFO  simpleserver.util.prop - Using poperty file: resources/simpleserver.properties
lein test simpleserver.testutils.users-util
lein test simpleserver.userdb.users-test
lein test simpleserver.webserver.server-test
lein test simpleserver.webserver.session-test
Ran 11 tests containing 47 assertions.
0 failures, 0 errors.
real	0m6.027s
user	0m7.632s
sys	0m0.383s
```

**Java**:

```bash
$ time ./gradlew --rerun-tasks test
Test result: SUCCESS
Test summary: 15 tests, 15 succeeded, 0 failed, 0 skipped
BUILD SUCCESSFUL in 5s
5 actionable tasks: 5 executed
real    0m5.757s
user    0m1.080s
sys	0m0.131s
```

**Javascript**:

```bash
time ./run-tests-with-trace.sh
  28 passing (94ms)
real	0m0.775s
user	0m0.683s
sys	0m0.083s
```

**Python**:

```bash
time ./run-pytest.sh 
========================================== test session starts 
platform linux -- Python 3.6.6, pytest-3.9.3, py-1.7.0, pluggy-0.8.0
rootdir: /mnt/edata/aw/kari/github/python/webstore-demo/simple-server, inifile: setup.cfg
collected 14 items                                                            
tests/domaindb/test_domain.py ....                                                                 [ 28%]
tests/userdb/test_users.py ...                                                                     [ 50%]
tests/webserver/test_server.py ......                                                              [ 92%]
tests/webserver/test_session.py .                                                                  [100%]

======================================= 14 passed in 0.11 seconds 
real	0m0.416s
user	0m0.376s
sys	0m0.039s
```

**Go**:

```bash
time ./run-go-tests.sh 
```

**The results are:**

| Language      |  Time  |
| ------------- |-------:|
| Clojure       |   6.0s |
| Java          |   5.8s |
| Javascript    |   0.8s |
| Python        |   0.4s |
| Go            |   TODO |

It's pretty obvious that Clojure and Java lose the contest because of the loading of JVM. TODO: Comment regarding Go.


# Go REPL

TODO


# Logging

My good friend and Go guru Tuomo Varis told me not to use external libraries but to do everything using Go standard library (to learn it better). I considered this a moment but first decided not to follow his good recommendation. The rationale being that I wanted to quickly implement the core functionalities of a web server and e.g. not to reinvent a logging framework myself. Therefore I used one of the most used Go logging framework [Logrus](https://github.com/sirupsen/logrus). 

But when discussing with Tuomo and he convinced me that implementing a simple logger based on the Go standard library logger should be rather simple I took the challenge and implemented my own custom [logger.go](https://github.com/karimarttila/go/blob/master/simpleserver/util/logger.go) based on the Go standard library logger. Basically I just implemented various log levels, my custom function entry/exit logging and some custom formatting of log entries.

Example of logging output:

```text
[2018-11-06T20:04:05.507Z] - [DEBUG] [main.main] - ENTER
[2018-11-06T20:04:05.507Z] - [DEBUG] [main.main] - Starting server...
[2018-11-06T20:04:05.507Z] - [DEBUG] [main.main] - - Port: 4047
[2018-11-06T20:04:05.507Z] - [DEBUG] [main.main] - - Report_caller: true
[2018-11-06T20:04:05.507Z] - [DEBUG] [main.main] - - Log_level: debug
[2018-11-06T20:04:05.508Z] - [DEBUG] [main.main] - - Log_file: src/github.com/karimarttila/go/simpleserver/logs/simpleserver.log
[2018-11-06T20:04:05.508Z] - [DEBUG] [webserver.StartServer] - ENTER
[2018-11-06T20:04:05.508Z] - [DEBUG] [webserver.handleRequests] - ENTER
[2018-11-06T20:04:06.673Z] - [DEBUG] [webserver.getInfo] - ENTER
[2018-11-06T20:04:06.673Z] - [DEBUG] [webserver.getInfo] - EXIT
```

Old a bit tweaked Logrus did the following output:

```text
time="2018-11-05T21:59:36+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.handleRequests debugtype=ENTER
time="2018-11-05T22:00:01+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.getInfo debugtype=ENTER
time="2018-11-05T22:00:01+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.getInfo debugtype=EXIT
```

So, you can decide which one is better. Personally I like my output better since it is more concise and readable (when implementing new server I always log at debug level all method/function entries/exits to provide good insight what is happening in the server while running in development mode - in production debug is turned off, of course). 

I also like the idea that there are no extra dependencies because of the logger but the logger is implemented using just Go standard library.

Thanks Tuomo for being stringent with my Go studies!


# Readability

TODO

Let's use Python and Go implementations as an examples of readability of those languages (you can check equivalent examples of Javascript, Java and Clojure in my previous blog posts, see links in the beginning of this article):


**Python**:

```python
def test_get_product_groups(client):
    myLogger.debug(ENTER)
    token = get_token(client)
    decoded_token = (b64encode(token.encode())).decode()
    mimetype = 'application/json'
    headers = {
        'Content-Type': mimetype,
        'Accept': mimetype,
        'Authorization': 'Basic ' + decoded_token
    }
    url = '/product-groups'
    response = client.get(url, headers=headers)
    assert response.status == '200 OK'
    json_data = json.loads(response.data)
    assert json_data.get('ret') == 'ok'
    assert b"product-groups" in response.data
    product_groups = json_data.get('product-groups')
    assert product_groups['1'] == 'Books'
    assert product_groups['2'] == 'Movies'
    myLogger.debug(EXIT)
```

**Go**:

```go
TODO
```

TODO: Comment.



# Productivity

TODO


# Lines of Code

Let's once again compare the lines of code between different implementations:


| Language      | Files  | LoC    |
| ------------- |-------:| ------:|
| Java          |     30 |   1612 |
| Javascript    |      7 |    674 |
| Clojure       |      6 |    612 |
| Python        |      8 |    528 |
| Go            |   TODO |  TODO  |


TODO: Comment.

# Performance

TODO



# Conclusions

TODO
