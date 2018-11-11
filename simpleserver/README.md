# Go Simple Server  <!-- omit in toc -->


# Table of Contents  <!-- omit in toc -->
- [Introduction](#introduction)
- [Learning Process](#learning-process)
- [Go](#go)
- [GoLand](#goland)
- [Code Format](#code-format)
- [Static Code Analysis](#static-code-analysis)
- [JSON Web Token](#json-web-token)
- [Error handling](#error-handling)
- [Go Interfaces](#go-interfaces)
- [Testing](#testing)
- [GoLand Debugger](#goland-debugger)
- [Map, Reduce and Filter](#map-reduce-and-filter)
- [Test Performance Between Five Languages](#test-performance-between-five-languages)
- [Go REPL](#go-repl)
- [Logging](#logging)
- [Readability](#readability)
- [Productivity](#productivity)
- [Lines of Code](#lines-of-code)
- [Performance](#performance)
- [Go vs Java](#go-vs-java)
- [Conclusions](#conclusions)


# Introduction

It's really fun and interesting to learn a new programming language. When you have zero knowledge before taking the challenge learning the new language is a kind of exploration expedition to a new land with different language and customs - the only way to effectively learn it is to go there, learn the language and its idioms.

So, Go (or 'golang' for search engine friendliness) is my fifth language I'm implementing the Simple Server. The previous versions are:

- [Clojure](https://github.com/karimarttila/clojure/tree/master/clj-ring-cljs-reagent-demo/simple-server)
- [Javascript/Node](https://github.com/karimarttila/javascript/tree/master/webstore-demo/simple-server)
- [Java](https://github.com/karimarttila/java/tree/master/webstore-demo/simple-server)
- [Python](https://github.com/karimarttila/python/tree/master/webstore-demo/simple-server)

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

I use [GoLand](https://www.jetbrains.com/go/) as my Go IDE.

I use [IntelliJ IDEA](https://www.jetbrains.com/idea/) for Java programming, [PyCharm](https://www.jetbrains.com/pycharm) for Python programming and IntelliJ IDEA with [Cursive](https://cursive-ide.com/) plugin for Clojure programming. Since GoLand, PyCharm and IDEA are provided by the same company (JetBrains) they provide very similar look-and-feel. So, there are a lot of synergy benefits to use the same IDE for several programming languages.

GoLand is really great for Go development. I created a test run configuration for each package and while I was developing that package I once in a while ran the equivalent GoLand test run configuration. If there were some errors it was very fast to add a debugger breakpoint to the failed test, hit the debugger and check the system state (variable values...) in the breakpoint. Go compiles extremely fast and the GoLand debugger starts blazingly fast so developing Go code like this was really fast. Using the GoLand debuggur it's also a nice way to look what's inside the standard library entities.


# Code Format

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

# JSON Web Token

Damn, I need one dependency, after all. I was hoping I could implement the Simple Server just using Go standard library but there is no JSON Web Token manipulation in the Go standard library and I really don't want to implement some poor JSON Web Token library myself for this project. So, I'm using [jwt-go](https://github.com/dgrijalva/jwt-go). Sorry, Tuomo. 

# Error handling

I kind of like Go's error handling. I have done production software with C some 20 years ago and you always had to be pretty carefull with returned error codes. In C++ you could define exceptions and throwing and catching them which kind of simplified error handling but also with a certain price. Java adopted the exception strategy but divided exceptions to runtime exceptions which you didn't have to explicitely handle and checked exceptions which you had to explicitely handle - many consider this a failed experiment since e.g. Spring exclusively uses just runtime exceptions. Go has a different strategy. There are no exceptions in the language but you can return many return values. An idiomatic way is to return the actual return value and an error - if there are no errors error value is nill, if there were errors the error value provides indication of the error. This is kind of nice but once again comes with a price - makes the error handling more explicit but creates more manual work for the programmer to handle errors.

Example:

```go
func CreateJsonWebToken(userEmail string) (ret string, err error) {
	util.LogEnter()
	expStr := util.MyConfig["json_web_token_expiration_as_seconds"]
	expiration, err := strconv.Atoi(expStr)
	if err != nil {
		util.LogError("Error converting json_web_token_expiration_as_seconds: " + expStr)
	} else {
		ttl := time.Duration(expiration) * time.Second
		claimExp := time.Now().UTC().Add(ttl).Unix()
		myClaim := SSClaim{
			userEmail,
			jwt.StandardClaims{
				ExpiresAt: int64(claimExp),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
		ret, err = token.SignedString(superSecret)
		if err != nil {
			util.LogError("error signing json web token: " + err.Error())
		} else {
			mySessions[ret] = true
		}
	}
	util.LogExit()
	return ret, err
}
... and the test class:
func TestJsonWebToken(t *testing.T) {
	util.LogEnter()
	testEmail := "kari.karttinen@foo.com"
	jsonWebToken, err := CreateJsonWebToken(testEmail)
	if err != nil {
		t.Error("CreateJsonWebToken returned error: " + err.Error())
	}
...
```

Some Go programmers immediately recognize that you could make the else branches mostly disappear just returning the error when you encounter one. I'm a bit of an old school programmer and I don't like to have many return points in my functions - I myself like the way in which the functionality goes to the end of the function and only there is the return point of the function.

Just out of curiosity, let's compare this to the version in which we return immediately when we encounter an error:

```go
func CreateJsonWebToken(userEmail string) (ret string, err error) {
	util.LogEnter()
	expStr := util.MyConfig["json_web_token_expiration_as_seconds"]
	expiration, err := strconv.Atoi(expStr)
	if err != nil {
		util.LogError("Error converting json_web_token_expiration_as_seconds: " + expStr)
		util.LogExit()		
		return ret, err
	}
	ttl := time.Duration(expiration) * time.Second
	claimExp := time.Now().UTC().Add(ttl).Unix()
	myClaim := SSClaim{
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: int64(claimExp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	ret, err = token.SignedString(superSecret)
	if err != nil {
		util.LogError("error signing json web token: " + err.Error())
		util.LogExit()
		return ret, err
	}
	mySessions[ret] = true
	util.LogExit()
	return ret, err
}
```

Ok. We got out of the else branches, but we introduced some extra mental burden when we have to remember that we now can exit from the function from many points. E.g. I like to see in my debug log every function entry and exit - now we have to remember to add the LogExit to every place we return from the function. Also this is a bad practice if you have allocated some resources in the beginning of the function and you need to deallocate the resources before exiting the function - now you have to remember to deallocate the resource in many points. So, I guess I keep my old habits. 

The latter version has 28 lines and my preferred version has only 26 lines - so after all we didn't even save any lines.

And what does say one of my programming gurus about this? Let's take McConnell's Code Complete ed. 2 from by book shelf and search the relevant chapter (17.1): "**Minimize the number of returns in each routine**. It’s harder to understand a routine when, reading it at the bottom, you’re unaware of the possibility that it returned somewhere above. For that reason, use returns judiciously - only when they improve readability." 

# Go Interfaces

Go interfaces are actually pretty nice minimalistic way to provide abstraction and reuse to Go code. Example in [server.go](https://github.com/karimarttila/go/blob/master/simpleserver/app/webserver/server.go):

I realized that various http API calls have a bit different error JSON structures returned and I was wondering that can I use Go interfaces to provide some abstraction to different structs. It turned out that I could:

```go
// Using ErrorResponder interface we can make error handling generic for all API calls.
type ErrorResponder interface {
	GetFlag() bool
	GetMsg() string
	WriteError(writer http.ResponseWriter) (err error)
}

// Used by all ErrorResponder entities.
func getEncoder(writer http.ResponseWriter) (encoder *json.Encoder) {
	encoder = json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	return encoder
}

// ErrorResponse is the base struct for all web layer Error response entities.
type ErrorResponse struct {
	Flag bool   `json:"-"` // Just to tell the whether we have initialized this struct or not (zero-value for bool is false, i.e. if the value is ready we know that we have initialized the struct).
	Ret  string `json:"ret"`
	Msg  string `json:"msg"`
}

func (e ErrorResponse) GetFlag() bool {
	return e.Flag
}

...
// SigninErrorResponse is the /signin API error response entity.
type SigninErrorResponse struct {
	ErrorResponse
	Email string `json:"email"`
}

func (e SigninErrorResponse) GetFlag() bool {
	return e.Flag
}
```

So, SigninErrorResponse extends the basic ErrorResponse, which provides basic error JSON fields. API says that SigninErrorResponse needs to provide also email field, so SigninErrorResponse is nice to implement in Go by extending ErrorResponse and adding the required email field. 

Then we ErrorResponder interface and both error response structs implements it. Then later in code we can:


```go
//in postSignin:
	if signinErrorResponse.Flag {
		writeError(writer, signinErrorResponse)
	}
//... and in postLogin:
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
// I.e. we can handle both error responses the same way, and the writeError function uses the interface ErrorResponnder and not the actual structs:
func writeError(writer http.ResponseWriter, errorResponder ErrorResponder) {
	util.LogEnter()
	err := errorResponder.WriteError(writer)
	if err != nil {
		// Everything else failed, just write the json as string to http.ResponseWriter.
		writer.Write([]byte(`{"ret":"failed","msg":"A total failure, original error: ` + errorResponder.GetMsg() + `"}`))
	}
	util.LogExit()
}
```

So, using this interface abstraction I could streamline the error handling pretty generic regardless of the struct different API calls use to report errors.



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


# Map, Reduce and Filter

There are no map, reduce and filter implementations in the Go standard library because Go is a statically typed language which does not provide generics - you either should have a dynamically typed language (like Clojure, Javascript or Python) or a statically typed language with generics (like Java) to have real map, reduce and filter functions. I googled this a bit and found one of Go's inventors, Rob Pike's [filter](https://github.com/robpike/filter) implementation in which he says: 

```text
"I wanted to see how hard it was to implement this sort of thing in Go, with as nice an API as I could manage. It wasn't hard. Having written it a couple of years ago, I haven't had occasion to use it once. Instead, I just use "for" loops. You shouldn't use it either."
```

So, let's just use for loops while programming Go. This is a bit of a pity since map, reduce and filter are very idiomatic e.g. in functional languages like Clojure. But you just have to accept that when in Rome do as the Romans do. 

A couple of examples how to implement getProduct using filter in Java, Python, Javascript and Clojure, and for loop in Go:

**Java**:

```java
List<Product> result = products.stream().filter(thisProduct ->
        (thisProduct.pId == pId) && (thisProduct.pgId == pgId))
        .collect(Collectors.toList());
// There should be 0 or 1.
if (result.size() == 1) {
    product = result.get(0);
}
else {
    logger.error("Didn't find exactly one product, count is: {}", result.size());
}
```

**Python**:

```python
product = list(filter((lambda x: x[0] == str(p_id)), raw_products))
ret = product[0] if (len(product) > 0) else None
```

**Javascript**:

```javascript
const filtered = rawProducts.filter(row => row[0] === `${pId}` && row[1] === `${pgId}`);
const product = filtered[0];
```

**Clojure**:

```clojure
(let [products (-get-raw-products pg-id)]
  (first (filter (fn [item]
                   (let [id (first item)]
                     (= id (str p-id))))
                 products))))
```

**Go**:

```go
for _, product := range rawProductsList {
	if product.PId == pId {
		ret = product
		break
	}
}
```

Once again: Java is verbose, Python and Javascript are concise, Clojure is elegant and Go is very C-like. I spent some time browsing those five implementations and I realized something. I have done so much production code in Java that error handling comes from the spine. For some languages in this exercise I didn't bother that much to do error handling (e.g. the Javascript implementation above in which we should test if filtered has exactly one item or not (as in the Java implementation)). Well, this was just an exercise. Maybe I'll do a code review for myself later on with all these five implementations and fix error handling in all of them.

Another interesting observation is that how the language drives the thinking in implementation. In other languages I have created this idea of raw products (all 8 fields per product, versus the actual product which has only the 4 fields needed when returning the product list), but for Java I have created just the Product class which has all 8 fields. Weird, I need to look into that later on when I have more time. All implementations provide the exact same API, though. I even used the Simple Frontend to test all Simple Frontend implementations that session handling works the same way and all pages (product groups, products list and product) look the same.

Maybe I'll also create a blog post later on in which I'll compare the differences in those five languages a bit deeper. Might be pretty interesting when all those implementations are fresh in my memory.


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

# Go vs Java

Go is so much better that Java in many ways:
- Go compiles to bare metal - Java compiles to byte code which you run in a JVM - overhead.
- Go's error handling is simpler.
- Go is much more concise than Java - code is easier to handle.
- In Go no class hell - in Java you realize that you are creating class this and class that all the time.

TODO: Examples here.




# Conclusions

I fell in love with Go. Go is really a very concise and productive language if you need a robust and performant statically typed language with excellent concurrency support. Much better than Java which compared to Go is verbose, non-productive and concurrency support is far behind Go. I have done quite a lot of C++ and Java programming and I must say that Go's error handling with idiomatic error entity as paired with the actual return value from functions is really great and simple. Go is definitely going to be my choice of statically typed language in my future projects. But still, if I need to create a quick script, e.g. a surrogate script for handling aws cli calls and process return json - I will choose Python. And if I need to process a lot of data - Clojure. But when you need statically typed language and excellent performance with great concurrency support - Go. 

I started my programming career with C. Hacking Go is a bit like coming home, except you don't need to be meticulous with memory allocation / deallocation. I think Go gives all goodies from C programming but takes care of the heavy lifting of what's difficult in C. Go code is really simple and elegant - the language provides the exact support for those things that you really need and doesn't add anything extra to the language (like Einstein put it: "Everything should be made as simple as possible, but not simpler").

The feeling was actually quite amazing. I started my Go hacking with practically zero Go knowledge on Monday, and on Saturday I felt like all pieces of the puzzle just locked in to the right places and creating code was really fluent and easy. 




TODO
