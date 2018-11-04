# Go Simple Server  <!-- omit in toc -->


# Table of Contents  <!-- omit in toc -->
- [Introduction](#introduction)
- [Tools and Versions](#tools-and-versions)
- [Go](#go)
- [GoLand](#goland)
- [Go Code Format](#go-code-format)
- [Testing](#testing)
- [Go REPL](#go-repl)
- [Logging](#logging)
- [Readability](#readability)
- [Productivity](#productivity)
- [Lines of Code](#lines-of-code)
- [Performance](#performance)
- [Conclusions](#conclusions)


# Introduction


# Tools and Versions


# Go

I was using [Go](https://golang.org//) 1.11 on Ubuntu18 when implementing this Simple Server.

You have to set the $GOPATH and $GOROOT environmental variables to point to your Go project directory and where your Go installation is. See example in (setenv.sh)[TODO].

```bash
go version      =>go version go1.11.1 linux/amd64
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



# Testing

TODO

Some performance aspects of tests:

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

TODO


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
