# Go Simple Server  <!-- omit in toc -->


# Table of Contents  <!-- omit in toc -->
- [Introduction](#introduction)
- [Tools and Versions](#tools-and-versions)
- [Go](#go)
- [GoLand](#goland)
- [Go Code Format](#go-code-format)
- [Testing](#testing)
- [Python REPL](#python-repl)
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

Run the tests in console:

```bash
./run-pytest.sh
```

The tests are implemented using [pytest](https://docs.pytest.org/en/latest/). Pytest is pretty straightforward to use and PyCharm also provides nice integration to unit tests implemented using pytest (debugger etc.). In PyCharm just create a new pytest configuration and configure it to use your specific pytest file and you are good to go to run that test and the code it calls in your debugger.

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


**The results are:**

| Language      |  Time  |
| ------------- |-------:|
| Clojure       |   6.0s |
| Java          |   5.8s |
| Javascript    |   0.8s |
| Python        |   0.4s |

It's pretty obvious that Clojure and Java lose the contest because of the loading of JVM. But I was surprised that Python runs the tests that fast.


# Python REPL

Python REPL is one of the best REPLs I have used outside the Lisp world. In the Lisp world REPLs are real REPLs which allow you to experiment with the live system in ways that no other REPL or debugger lets you do in other languages - it's pretty impossible to explain this, you just have to learn some Lisp and try yourself (e.g. [Clojure](https://clojure.org/)). So, now that we have gone through my mandatory Clojure advertisement let's go back to Python REPL. Compared to JShell Java REPL Python wins the fight hands down. Python REPL came with the first version of Python (we just had to wait some 20 years for Java REPL) and because Python is dynamically typed language the REPL is pretty easy to use (compared to Java JShell which is really awkward to use, even with a good IDE). 

PyCharm provides a nice REPL, an example follows:

```python
 >>> runfile('/mnt/edata/aw/kari/github/python/webstore-demo/simple-server/simpleserver/domaindb/domain.py', wdir='/mnt/edata/aw/kari/github/python/webstore-demo/simple-server')
>>> myD = Domain()
2018-10-30 18:40:11,769 - __main__ - __init_product_db - DEBUG - ENTER
2018-10-30 18:40:11,770 - __main__ - __read_product_groups - DEBUG - ENTER
...
2018-10-30 18:40:11,771 - __main__ - __read_raw_products - DEBUG - EXIT
2018-10-30 18:40:11,771 - __main__ - __init_product_db - DEBUG - EXIT
>>> myD.get_raw_products(1)
[['2001', '1', 'Kalevala', '3.95', 'Elias Lönnrot', '1835', 'Finland', 'Finnish'], ...]
```

So, using the runfile method you are able to reload any module to Python console and then try the methods there in isolation.


# Logging

What a relief Python logging configuration is after Spring hassle. You just create the [logging.conf](https://github.com/karimarttila/python/blob/master/webstore-demo/simple-server/logging.conf) file and that's about it. 


# Readability

Python wins Javascript and pretty much any language in readability hands down. It is probably the most readable language there is. I would say that it is even more readable than Clojure which is also a very readable language (once you get used to read a functional language). Java loses to Python in readability just because of the monstrous verbosity of the language.

Let's use Javascript and Python implementations as an examples of readability of those languages (you can check equivalent examples of Java and Clojure in my previous blog posts, see links in the beginning of this article):

**Javascript**:

```javascript
  describe('GET /product-groups', function () {
    let jwt;
    it('Get Json web token', async () => {
      // Async example in which we wait for the Promise to be
      // ready (that i.e. the post to get jwt has been processed).
      const jsonWebToken = await getJsonWebToken();
      logger.trace('Got jsonWebToken: ', jsonWebToken);
      assert.equal(jsonWebToken.length > 20, true);
      jwt = jsonWebToken;
    });
    it('Successful GET: /product-groups', function (done) {
      logger.trace('Using jwt: ', jwt);
      const authStr = createAuthStr(jwt);
      supertest(webServer)
        .get('/product-groups')
        .set('Accept', 'application/json')
        .set('Authorization', authStr)
        .expect('Content-Type', /json/)
        .expect(200, {
          ret: 'ok',
          'product-groups': { 1: 'Books', 2: 'Movies' }
        }, done);
    });
  });
```

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

I would say that Python is more readable.


# Productivity

What a joy it was to program Python after Java. Dynamically typed language! Concise! Clear syntax! Simple! The productivity of Python programming compared to Java is like from another planet - I explored PyCharm new features and Flask the first evening and in the second evening I implemented the domaindb module and most of the userdb module and related unit tests. The third evening I implemented the rest of the application and the unit tests and that was it. There were absolutely no hassle in Python code, configurations or in IDE (PyCharm). 

Python (and especially PyCharm) **REPL** is definitely the best REPL I have used outside Lisp world. The Python REPL makes exploring small code snippets a breeze. 

Using PyCharm debugger is also so easy and fast. If you have even minor issues in your code you tend to add a breakpoint and hit the debugger. This is actually pretty interesting since in the Lisp world you hardly ever use the debugger - you tend to have a live REPL to your system while you add new functionalities to the system. You can't have a live REPL to your Python system in the same sense but PyCharm debugger is a pretty good second option. And when you compare Python debugger to Java debugger - Python is lightning fast to start. Creating Run configurations for your unit tests in PyCharm is also very easy and straightforward. PyCharm debugger is also a great tool to check what's inside various entities (e.g. I just used the debugger to check where the http status code is inside the Flask response entity and what is its name) - if you are lazy to search that information in thelibrary API documentation.

In general I think Python must be the most productive language I have ever used. Clojure might win the case in productivity after a couple of years of serious Clojure hacking but Python is unbeatable in the scripting category - you may have months of gaps between your Python hacking sessions but the language is always easy to put into real work regardless how long it was you programmed Python the last time.

If you compare Python to Java - Python wins hands down. Java is verbose - Python is concise. Java has long development cycle (edit, compile, build, load to JVM, run) - Python has short development cycle (edit, run). Java has difficult syntax - Python has very easy syntax.

If you compare Python to Javascript/Node, Python wins in clean syntax and overall easyness to create and test code. 

There is nothing inherently bad in Python. I would be cautious to use Python in a big project with tens of developers working in the same code base unless you have some strict rules how to protect collaboration from the typical mistakes using dynamically typed language in a big project (e.g. mandatory use of [type hints](https://docs.python.org/3/library/typing.html)). 


# Lines of Code

Let's once again compare the lines of code between different implementations:


| Language      | Files  | LoC    |
| ------------- |-------:| ------:|
| Java          |     30 |   1612 |
| Javascript    |      7 |    674 |
| Clojure       |      6 |    612 |
| Python        |      8 |    528 |

If you drop the empty package files (```__init__.py```) there are only 8 source code files in the production source tree and altogether only 582 lines of code. So, it seems that Python is the winner of this part of the contest.


# Performance

The [GIL](https://wiki.python.org/moin/GlobalInterpreterLock) might cause some issues if you try to create a system which should be responsive to a large amount of events / sessions. Node is also single-threaded but Node has a special architecture in which Node runs single thread in an event loop and delegates e.g. I/O work to worker pool threads. This makes Node extremely efficient in handling tasks which are not CPU intensive (on the other side CPU intensive tasks may degrade the Node performance quite a lot). Java system on the contrary typically spins a dedicated thread for each request. This is more expensive (consumes more machine resources) but one thread (for one client) does not block processing of another thread (client). Python has the infamous Global Interpreter Lock which has generated a lot of debate in the Python community during Python's lifetime. In most cases this is not a problem since you usually use Python for small tasks. But if you use Python for CPU intensive work handling a huge set of tasks or requests in parallel you have to find some special solutions for it (and those do exist if you google them, see e.g. ["Efficiently Exploiting Multiple Cores with Python"](http://python-notes.curiousefficiency.org/en/latest/python3/multicore_python.html)).


# Conclusions

I have used Python for some 20 years for various tasks. E.g. my personal backup scripts at home are implemented using Python. My virtual dog was implemented using Python (watched my IP cameras and started an audio file of an insane dog barking if any movement detected at my backyard - I can reveal this now since I have a real big insane dog watching the house). At work I have used Python for various log analysis work, watching processes, analyzing copy-pasting of Java classes between projects, gluing various aws cli commands together and so on and so on. Python is really a good language for quick ad hoc scripts you may need for various purposes. Python is also pretty good language to implement a web server as I now did (performance testing is to be done). 

This exercise to implement the same web server using five languages (one still to go) has been a real eye opener. I must say that I'm pretty disappointed regarding Java's productivity. I thought that after 20 years I could have implemented the web server faster than with a totally new language (Javascript), but it took the same amount of time (about 3 weeks) for both of them. Python - 3 evenings. Now that those big enterprise monolithic systems are going to be history and the new era of Cloud serveless and microservices are emerging one can choose the language much freely. I think I leave Java in that big monolithic enterprise world where it belongs. The new serverless implementations are best to implement using Python or Javascript. Data oriented microservices possibly using Clojure. 

The last challenge is still ahead: [Go](https://golang.org/). It's going to be pretty interesting since I know nothing of Go. So, one more time I will raise the old Simple Server ghost and mold it into an implementation - this time using Go.