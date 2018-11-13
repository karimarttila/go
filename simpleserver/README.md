# Go Simple Server  <!-- omit in toc -->


# Table of Contents  <!-- omit in toc -->
- [Introduction](#introduction)
- [Learning Process](#learning-process)
- [Go](#go)
- [GoLand](#goland)
- [Code Format](#code-format)
- [Static Code Analysis](#static-code-analysis)
- [JSON Web Token](#json-web-token)
- [Http and REST Handling](#http-and-rest-handling)
- [Error Handling](#error-handling)
- [Go Interfaces](#go-interfaces)
- [Testing](#testing)
- [GoLand Debugger](#goland-debugger)
- [Map, Reduce and Filter](#map-reduce-and-filter)
- [Test Performance Between Five Languages](#test-performance-between-five-languages)
- [Go Playground](#go-playground)
- [Logging](#logging)
- [Readability](#readability)
- [CORS](#cors)
- [Productivity](#productivity)
- [Lines of Code](#lines-of-code)
- [Performance](#performance)
- [Concurrency Support](#concurrency-support)
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

I had practically zero knowledge of Go when I started this project. I watched one short "Go Basics" type Pluralsight video before starting to program Go. Mostly everything I just learned on the fly while doing programming. E.g. I was wondering what to do in a situation in which some Go function returns three distinct return values but I need only one? I handled these situations just by googling, e.g.  ```go function returns multiple values "unused variable"``` => use underscore for those variables you don't need.

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

You have to set the $GOPATH and $GOROOT environmental variables to point to your Go project directory and where your Go installation is. See example in [setenv.sh](https://github.com/karimarttila/go/blob/master/simpleserver/setenv.sh).

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

This is  a one time task and creates the [Gopkg.toml](https://github.com/karimarttila/go/blob/master/simpleserver/Gopkg.toml) file.

Running ```dep ensure``` in the project (simpleserver) updates the [Gopkg.lock](https://github.com/karimarttila/go/blob/master/simpleserver/Gopkg.lock) file and creates a vendor directory (not in Git). 

I created a Git repo for this project and added it under the GOPATH in my machine. I also used the go get <url> to load the source code of the dependencies that I'm using (just "github.com/dgrijalva/jwt-go" library for handling JSON web token).

And so we finalized our short tour to "Go and package management."

BTW. There has been a lot of criticism regarding Go's package management (being sources in Github). Package management may change in some future Go version.


# GoLand

I used [GoLand](https://www.jetbrains.com/go/) as my Go IDE.

I use [IntelliJ IDEA](https://www.jetbrains.com/idea/) for Java programming, [PyCharm](https://www.jetbrains.com/pycharm) for Python programming and IntelliJ IDEA with [Cursive](https://cursive-ide.com/) plugin for Clojure programming. Since GoLand, PyCharm and IDEA are provided by the same company (JetBrains) they provide very similar look-and-feel. So, there are a lot of synergy benefits to use the same IDE for several programming languages.

GoLand is really great for Go development. I created a test run configuration for each package and while I was developing that package I once in a while ran the equivalent GoLand test run configuration. If there were some errors it was very fast to add a debugger breakpoint to the failed test, hit the debugger and check the system state (variable values...) in the breakpoint. Go compiles extremely fast and the GoLand debugger starts blazingly fast so developing Go code like this was really fast. Using the GoLand debuggur it's also a nice way to look what's inside the standard library entities.


# Code Format

Go is an interesting language in that sense that formatting of the Go code is very opinionated. Very opinionated in that sense that the Go compiler even requests code to be in certain format or it doesn't compile the code even though it would be syntactically right. Formatting the Go code is build into the language (see: [format](https://golang.org/pkg/go/format/)).

You can reformat the Simple Server code using the following command in the $GOPATH directory.

```bash
go fmt github.com/karimarttila/go/simpleserver/app/...
```

I also provided a script [go-fmt-simpleserver.sh](https://github.com/karimarttila/go/tree/master/simpleserver/scripts) to run the fmt code formatter for all Simple Server Go files. It's a good idea to run this script every once in a while to keep the Go code formatting clean.


# Static Code Analysis

Go provides a simple static code analysis tool in the standard Go toolbox: [vet](https://golang.org/cmd/vet/). You can run the static code analysis using command: 

```bash
go vet github.com/karimarttila/go/simpleserver/...
```

I also found another interesting tool: [Staticcheck](https://staticcheck.io/docs/). Install the package and you can run staticcheck, gosimple and unused in one go for all Go code files:

```bash
megacheck github.com/karimarttila/go/simpleserver/app/...
```

Staticcheck open source version is free. If you find the tool useful you should consider buying the commercial version. 


# JSON Web Token

Damn, I needed one dependency, after all. I was hoping I could implement the Simple Server just using Go standard library but there is no JSON Web Token manipulation in the Go standard library and I really don't want to implement some poor JSON Web Token library myself for this project. So, I'm using [jwt-go](https://github.com/dgrijalva/jwt-go). Sorry, Tuomo. 


# Http and REST Handling

Go is an unusual language in that sense that it provides excellent standard library and therefore you seldom need to import extra dependencies. There are some external http routing libraries in Go, e.g. [Gorilla](https://github.com/gorilla/mux) but why to introduce external dependencies to your project if you can manage with the standard library? So, for http / REST handling I just used the Go standard library [net/http](https://golang.org/pkg/net/http/). It was pretty straightforward to use net/http, example below.

```go
func getProductGroups(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	writeHeaders(writer)
	if request.Method == "OPTIONS" {
		return
	}
	parsedEmail, errorResponse := isValidToken(request)
	var productGroups domaindb.ProductGroups
	if !errorResponse.Flag {
		util.LogTrace("parsedEmail from token: " + parsedEmail)
		productGroups = domaindb.GetProductGroups()
		encoder := json.NewEncoder(writer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(productGroups)
		if err != nil {
			errorResponse = createErrorResponse(err.Error())
		}
	}
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}
// ... and registering the API:
// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/signin", postSignin)
	http.HandleFunc("/login", postLogin)
	http.HandleFunc("/product-groups", getProductGroups)
	http.HandleFunc("/products/", getProducts)
	http.HandleFunc("/product/", getProduct)
	http.Handle("/", http.FileServer(http.Dir("./src/github.com/karimarttila/go/simpleserver/static")))
	log.Fatal(http.ListenAndServe(":"+util.MyConfig["port"], nil))
	util.LogExit()
}
```

If you wonder why the return is for OPTIONS - it's for CORS / preflight. 

# Error Handling

I kind of like Go's error handling. I have done production software with C some 20 years ago and you always had to be pretty carefull with returned error codes. In C++ you could define exceptions and throwing and catching them which kind of simplified error handling but also with a certain price. Java adopted the exception strategy but divided exceptions to runtime exceptions which you didn't have to explicitly handle and checked exceptions which you had to explicitly handle - many consider this a failed experiment since e.g. Spring exclusively uses just runtime exceptions. Go has a different strategy. There are no exceptions in the language (well, there is one exception) but you can return many return values. An idiomatic way is to return the actual return value and an error - if there are no errors then the error value is nill, if there were errors the error value provides indication of the error. This is kind of nice but once again comes with a price - makes the error handling more explicit but creates more manual work for the programmer to handle errors.

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
	if jsonWebToken == "" {
		t.Error("jsonWebToken is empty")
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

And what does one of my programming gurus say about this? Let's take McConnell's Code Complete ed. 2 from by book shelf and search the relevant chapter (17.1): "**Minimize the number of returns in each routine**. It’s harder to understand a routine when, reading it at the bottom, you’re unaware of the possibility that it returned somewhere above. For that reason, use returns judiciously - only when they improve readability." 

I implemented an interface based error handling (more about it in the next chapter). It made error handling pretty straigtforward and simple in API handlers, example:

```go
func postLogin(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	writeHeaders(writer)
	if request.Method == "OPTIONS" {
		return
	}
	var errorResponse ErrorResponse // Generic ErrorResponse will do for /login just fine.
	var loginData LoginData
	var loginResponse LoginResponse
	var jsonWebToken string
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		errorResponse = createErrorResponse("Decoding request body failed")
	} else {
		if loginData.Email == "" || loginData.Password == "" {
			errorResponse = createErrorResponse("Validation failed - some fields were empty")
		} else {
			credentialsOk := userdb.CheckCredentials(loginData.Email, loginData.Password)
			if !credentialsOk {
				errorResponse = createErrorResponse("Credentials are not good - either email or password is not correct")
			} else {
				jsonWebToken, err = CreateJsonWebToken(loginData.Email)
				if err != nil {
					errorResponse = createErrorResponse("Couldn't create token: " + err.Error())
				} else {
					loginResponse = LoginResponse{true, "ok", "Credentials ok", jsonWebToken}
					encoder := json.NewEncoder(writer)
					encoder.SetEscapeHTML(false)
					err := encoder.Encode(loginResponse)
					if err != nil {
						errorResponse = createErrorResponse(err.Error())
					}
				}
			}
		}
	}
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}
```

So, as you can see every step checks if there was an error, and if was, we just create an ErrorResponse and that's it. If we get to the end we know that everything went smoothly and we can write the actual payload to the http ResponseWriter. At the end we check if the errorResponse flag was turned on, and if we see this, we call writeError which according to interface writes the error response to the http ResponseWriter. I think this is clean enough for me. And if I compare it to Java's exception handling this is not bad at all. 


# Go Interfaces

Go interfaces are actually pretty nice minimalistic way to provide abstraction and reuse to Go code. Example in [server.go](https://github.com/karimarttila/go/blob/master/simpleserver/app/webserver/server.go):

I realized that various http API calls have a bit different error JSON structures returned and I was wondering if I can use Go interfaces to provide some abstraction to different structs. It turned out that I could:

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

func (e ErrorResponse) WriteError(writer http.ResponseWriter) (err error) {
	encoder := getEncoder(writer)
	err = encoder.Encode(e)
	return err
}

// SigninErrorResponse is the /signin API error response entity.
type SigninErrorResponse struct {
	ErrorResponse
	Email string `json:"email"`
}

func (e SigninErrorResponse) WriteError(writer http.ResponseWriter) (err error) {
	encoder := getEncoder(writer)
	err = encoder.Encode(e)
	return err
}
```

So, SigninErrorResponse extends the basic ErrorResponse, which provides basic error JSON fields. API says that SigninErrorResponse needs to provide also email field, so SigninErrorResponse is nice to implement in Go by extending ErrorResponse and adding the required email field. 

Then the ErrorResponder interface defined WriteError function and both error response structs implements the WriteError function. Then later in code we can:


```go
//in postSignin:
	if signinErrorResponse.Flag {
		writeError(writer, signinErrorResponse)
	}
//... and in postLogin:
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
// I.e. we can handle both error responses the same way, and the writeError function uses the interface ErrorResponder and not the actual structs:
func writeError(writer http.ResponseWriter, errorResponder ErrorResponder) {
	util.LogEnter()
	// NOTE: StatusOK is implicitely written first time writer.Write is called
	// unless other status code set.
	writer.WriteHeader(http.StatusBadRequest)
	err := errorResponder.WriteError(writer)
	if err != nil {
		// Everything else failed, just write the json as string to http.ResponseWriter.
		writer.Write([]byte(`{"ret":"failed","msg":"A total failure, original error: ` + errorResponder.GetMsg() + `"}`))
	}
	util.LogExit()
}
```

So, using this interface abstraction I could streamline the error handling pretty generic regardless of the struct different API calls use to report errors. Of course using Java the implementation would have been more generic since Java provides better abstractions in the language but I think Go's abstractions are good enough and do not bloat the language (a bit like Java does).



# Testing

Go is pretty amazing in that sense that you have also [the Go testing framework](https://golang.org/pkg/testing/) in the standard library.

Go tests are pretty easy to create. You don't have asserts but instead you just write standard application logic to test whether your package works as expected. Example:

```go
func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	// We could implement get this by querying /login, but let's make a shortcut.
	token, err := CreateJsonWebToken("kari.karttinen@foo.com")
	if err != nil {
		t.Errorf("Failed to get test token: %s", err.Error())
	}
	util.LogTrace("Test token: " + token)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	if err != nil {
		t.Errorf("Failed to base64 decode token: %s", err.Error())
	}
	//NOTE: We actually call directly the handler.
	// See below: "http.HandlerFunc(getProductGroups)...."
	request := httptest.NewRequest("GET", "http://localhost:"+port+"/product-groups", nil)
	request.Header.Add("authorization", "Basic "+encoded)
	recorder := httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(getProductGroups).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getProductGroups handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	response := recorder.Body.String()
	if len(response) == 0 {
		t.Error("Response was nil or empty")
	}
	// NOTE: Might look a bit weird, but it's pretty straightforward:
	// pgMap is a map (key:string), and values are maps, which keys are strings and values are strings.
	pgMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(response), &pgMap)
	if err != nil {
		t.Errorf("Unmarshalling response failed: %s", err.Error())
	}
	pg, ok := pgMap["product-groups"]
	if !ok {
		t.Errorf("Didn't find 'product-groups' in response")
	}
	pg1, ok := pg["1"]
	if !ok {
		t.Errorf("Didn't find product group 1 in response")
	}
	if pg1 != "Books" {
		t.Errorf("Product group 1 should have been 'Books'")
	}
	util.LogEnter()
}
```

The tests tend to be rather verbose, though. Compare to the equivalent Clojure test:

```clojure
(deftest get-product-groups-test
  (log/trace "ENTER get-product-groups-test")
  (testing "GET: /product-groups"
    (let [req-body {:email "kari.karttinen@foo.com", :password "Kari"}
          login-ret (-call-request ws/app-routes "/login" :post nil req-body)
          dummy (log/trace (str "Got login-ret: " login-ret))
          login-body (:body login-ret)
          json-web-token (:json-web-token login-body)
          params (-create-basic-authentication json-web-token)
          get-ret (-call-request ws/app-routes "/product-groups" :get params nil)
          dummy (log/trace (str "Got body: " get-ret))
          status (:status get-ret)
          body (:body get-ret)
          right-body {:ret :ok, :product-groups {"1" "Books", "2" "Movies"}}
          ]
      (is (not (nil? json-web-token)))
      (is (= status 200))
      (is (= body right-body)))))
```

The reason is of course that in Clojure we have a dynamically typed language which is designed for data processing (data is more data in code and more easily to be manipulated). 

You can run all tests in command line using command:

```bash
go test github.com/karimarttila/go/simpleserver/app/...
ok  	github.com/karimarttila/go/simpleserver/app/domaindb	(cached)
ok  	github.com/karimarttila/go/simpleserver/app/main	(cached)
ok  	github.com/karimarttila/go/simpleserver/app/userdb	(cached)
ok  	github.com/karimarttila/go/simpleserver/app/util	(cached)
ok  	github.com/karimarttila/go/simpleserver/app/webserver	(cached)
```

Running tests is pretty nice since Go compiles really fast and starts the tests immediately. Also if Go notices that the files have not changed the test results are cached. More about test performance in later chapter.


# GoLand Debugger

GoLand debugger is really good. Debugger starts immediately and is really fast. It's not a Lisp REPL but a pretty good second option. Go's statically typed data structures are pretty simple and GoLand debugger does a very good job presenting the data structures and values in the editor and in the variables window.


# Map, Reduce and Filter

There are no map, reduce and filter implementations in the Go standard library because Go is a statically typed language which does not provide generics - you either should have a dynamically typed language (like Clojure, Javascript or Python) or a statically typed language with generics (like Java) to have real map, reduce and filter functions. I googled this a bit and found one of Go's inventors, Rob Pike's [filter](https://github.com/robpike/filter) implementation in which he says: 

```text
"I wanted to see how hard it was to implement this sort of thing in Go, with as nice an API as I could manage. It wasn't hard. Having written it a couple of years ago, I haven't had occasion to use it once. Instead, I just use "for" loops. You shouldn't use it either."
```

So, let's just use for loops while programming Go. This is a bit of a pity since map, reduce and filter are a good abstraction and very idiomatic e.g. in functional languages like Clojure. But you just have to accept that when in Rome do as the Romans do. 

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
		if product[0] == wantedPid {
			found = product
			break
		}
	}
```

Once again: Java is verbose, Python and Javascript are concise, Clojure is elegant and Go is very C-like. I spent some time browsing those five implementations and I realized something. I have done so much production code in Java that error handling comes from the spine. For some languages in this exercise I didn't bother that much to do error handling (e.g. the Javascript implementation above in which we should test if filtered has exactly one item or not (as in the Java implementation)). Well, this was just an exercise. Maybe I'll do a code review for myself later on with all these five implementations and fix error handling in all of them.

Another interesting observation is that how the language drives the thinking in implementation. In other languages I have created this idea of raw products (all 8 fields per product, versus the actual product which has only the 4 fields needed when returning the product list), but for Java I have created just the Product class which has all 8 fields. Weird, I need to look into that later on when I have more time. All implementations provide the exact same API, though. I used the Simple Frontend to test all Simple Server implementations that session handling and API return values work the same way and all pages (product groups, products list and product) look the same.

Maybe I'll also create a blog post later on in which I'll compare the differences in those five languages a bit deeper. Might be pretty interesting when all those implementations are fresh in my memory.


# Test Performance Between Five Languages

All right! Finally my journey travelling through the five language lands is over and I can compare the languages. Let's first compare the test performance between the languages.


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
time GOCACHE=off scripts/go-test-simpleserver.sh
ok  	github.com/karimarttila/go/simpleserver/app/domaindb	0.004s
ok  	github.com/karimarttila/go/simpleserver/app/main	0.003s
ok  	github.com/karimarttila/go/simpleserver/app/userdb	0.008s
ok  	github.com/karimarttila/go/simpleserver/app/util	0.005s
ok  	github.com/karimarttila/go/simpleserver/app/webserver	0.007s
real	0m1.918s
```

**The results are:**

| Language      |  Time  |
| ------------- |-------:|
| Clojure       |   6.0s |
| Java          |   5.8s |
| Javascript    |   0.8s |
| Python        |   0.4s |
| Go            |   1.9s |

It's pretty obvious that Clojure and Java lose the contest because of the loading of JVM. Javascript and Python are pretty fast since they just start running the tests and hope that while interpreting the code there are no runtime errors. Go is statically compiled language and needs to compile first before running tests. So, Go's 2 seconds v. Java's 6 seconds is pretty good.


# Go Playground

There is no REPL in Go (very difficult to make for a statically typed language - it took some 20 years before we got some sort of very simple REPL for Java).

But there is some sort of workaround: [Go Playground](https://play.golang.org/). Gophers have created various examples in the Playground and you can try to google them. Example: [How to use map of maps](https://play.golang.org/p/pQsoB02pDl).


# Logging

My good friend and Go guru Tuomo Varis told me not to use external libraries but to do everything using Go standard library (to learn it better). I considered this a moment but first decided not to follow his good recommendation. The rationale being that I wanted to quickly implement the core functionalities of a web server and e.g. not to reinvent a logging framework myself. Therefore I first started to use one of the most used Go logging frameworks, [Logrus](https://github.com/sirupsen/logrus). 

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

A bit tweaked Logrus did the following output:

```text
time="2018-11-05T21:59:36+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.handleRequests debugtype=ENTER
time="2018-11-05T22:00:01+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.getInfo debugtype=ENTER
time="2018-11-05T22:00:01+02:00" level=debug msg="[]" caller=github.com/karimarttila/go/simpleserver/webserver.getInfo debugtype=EXIT
```

So, you can decide which one is better. Personally I like my output better since it is more concise and readable (when implementing new server I always log at debug level all method/function entries/exits to provide good insight what is happening in the server while running in development mode - in production debug is turned off, of course). 

I also like the idea that there are no extra dependencies because of the logger but the logger is implemented using just Go standard library.

Thanks Tuomo for being stringent with my Go studies!


# Readability

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
func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	// We could implement get this by querying /login, but let's make a shortcut.
	token, err := CreateJsonWebToken("kari.karttinen@foo.com")
	if err != nil {
		t.Errorf("Failed to get test token: %s", err.Error())
	}
	util.LogTrace("Test token: " + token)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	if err != nil {
		t.Errorf("Failed to base64 decode token: %s", err.Error())
	}
	//NOTE: We actually call directly the handler.
	// See below: "http.HandlerFunc(getProductGroups)...."
	request := httptest.NewRequest("GET", "http://localhost:"+port+"/product-groups", nil)
	request.Header.Add("authorization", "Basic "+encoded)
	recorder := httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(getProductGroups).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getProductGroups handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	response := recorder.Body.String()
	if len(response) == 0 {
		t.Error("Response was nil or empty")
	}
	// NOTE: Might look a bit weird, but it's pretty straightforward:
	// pgMap is a map (key:string), and values are maps, which keys are strings and values are strings.
	pgMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(response), &pgMap)
	if err != nil {
		t.Errorf("Unmarshalling response failed: %s", err.Error())
	}
	pg, ok := pgMap["product-groups"]
	if !ok {
		t.Errorf("Didn't find 'product-groups' in response")
	}
	pg1, ok := pg["1"]
	if !ok {
		t.Errorf("Didn't find product group 1 in response")
	}
	if pg1 != "Books" {
		t.Errorf("Product group 1 should have been 'Books'")
	}
	util.LogEnter()
}
```

Well, Python wins hands down. If you prefer readability and you don't care about performance and concurrency support - use Python.


# CORS

Damn [Cross-origin resource sharing](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) hit me again when testing with [Simple Frontend](https://github.com/karimarttila/clojure/tree/master/clj-ring-cljs-reagent-demo/simple-frontend). I spent some time figuring out how to provide simple configuration using the standard [net/http](https://golang.org/pkg/net/http/) library but didn't want to spend too much time with this since I wanted to finish this project this evening (and actually I solved it about 11pm but didn't have time to write the related blog post anymore). So, I tried [rs/cors](https://github.com/rs/cors) but I couldn't make it work properly. So turning back to standard library and reading documentation regarding CORS and preflight settings. Finally it was pretty easy to configure. I just had to add some headers for CORS and return from functions if method is OPTION (preflight). See details in [server.go](https://github.com/karimarttila/go/blob/master/simpleserver/app/webserver/server.go). I tested Go Simple Server with [Simple Frontend](https://github.com/karimarttila/clojure/tree/master/clj-ring-cljs-reagent-demo/simple-frontend) and now it works as previous 4 Simple Server implementations. 


# Productivity

Go productivity is not as good as in Python (Simple Server implementation took some 3 evenings in Python) and in Clojure (I would now do it in three evenings using Clojure), but Go productivity is much better than in Java (took some 3 weeks in Java even though I have programmed some 20 years of Java) (I estimated I had about one evening every week when I didn't do the implementation) and Javascript (took me also some 3 weeks and I had to learn the language and Node and all related tools while implementing Simple Server). Now about a week has passed and I'm done. So, Go's productivity seems to be somewhere between these languages. Let's summarize the rough results in a table to visualize them better.


| Language      |  Evenings | Experience before implementation (+ comments)   |
| ------------- |----------:|-------------------------------------------------|
| Java          |   18      | 20 years (I.e. 3 weeks * 6 evenings = 18)       |
| Go            |   8       | Zero                                            |
| Javascript    |   18      | Javascript some weeks, Node zero                |
| Clojure       |   3       | About one year. (I.e. 3 with current knowledge) |
| Python        |   3       | 20 years                                        |


So, I estimate that if I had at least a couple of years of experience of each language the table would look something like:

| Language      |  Evenings | MPI |
| ------------- |----------:|----:|
| Java          |   12      |   4 |
| Go            |   6       |   2 |
| Javascript    |   9       |   3 |
| Clojure       |   3       |   1 |
| Python        |   3       |   1 |


I dropped Java to 12 evenings since when implementing the Simple Server in Java I hadn't done serious Java programming for some 1,5 years and I spent quite a lot of time exploring new Java 10, new Spring functionalities, new IDEA features, how to use Java REPL etc. I could drop some 2 evenings away from Go if I had a couple of years of experience using it but Go and especially testing in Go is still a bit verbose which costs a couple of extra evenings compared to Python and Clojure (but Go's verbosity is nothing compared to Java's verbosity). Learning Javascript and Node took considerably more than learning Go (for both languages I started basically from Zero, but did Go implementation in about 8 evenings and Node implementation in about 18 evenings). I'm not that sure about Javascript's productivity but I would estimate it is about a bit higher than Go's productivity based on my experiences while learning the languages (but still considerably less than Java). Go's productivity is after all pretty good considering that it is a statically typed language. 

So, the clear winners of the productivity game are Python and Clojure. Therefore: my final recommendations for implementing backend systems based on my experiences of these languages are:

- **Python**: If you need to implement simple scripts, surrogates to aws cli etc: use Python. Python provides also excellent libraries for ML and mathematical libraries are actually implemented in C (and Python just provides frontend to the libraries) - so in ML Python is also pretty performant.
- **Clojure**: If you need to manipulate a lot of data and you need excellent concurrency support: use Clojure. The only downside with Clojure is the high learning curve for new developers (since the functional language can be rather mind blowing) - therefore the developer pool is going to be always much scarcer than in other languages.
- **Go**: If you need bare metal performance with excellent concurrency support and you don't need to manipulate a lot of data: use Go. Go probably is best as a system tool language - data manipulation is a lot more verbose than in Python and Clojure.
- **Java**: If you have a really big enterprise system and and offshore development with tens of developers working with the same code base, probably Java.
- **Javascript**: And finally Javascript if you specifically for some reason need to use Node (e.g. the team consists of younger developers who only know Javascript - this is quite common, actually). 

I have nothing against Javascript/Node - implementing the Simple Server using Node was pretty straightforward and the tools were good. But I really couldn't find other good reasons to use Node since regarding various aspects there almost always is a better backend language.

I also created a new index - Marttila Productivity Index for Programming Languages (**MPI**) (with tongue in cheek). We divide all results of all languages with the most effective evenings (Python/Clojure: 3), and this way we get a relative productivy index for all languages (1 being the optimal number, the higher the number is from 1 the longer it takes to implement a web server using that language). So, the next time I start a new project and if I'm able freely to choose my language, I'll check my own recommendations and the MPI value I have given to each language. :-) 


# Lines of Code

Let's once again compare the lines of code between different implementations.

Headers:

- Language: language
- P-Files: production files (i.e. not including test files)
- P-LoC: production files total lines of code (i.e. not including test files)
- T-Files: test files
- T-LoC: test files total lines of code
- A-F: all files together
- A-LoC: all lines of code together
- MLCI: Marttila Lines of Code Index (divide all language's A-F with the lowest A-F)
- MPI: Marttila Productivity Index (from previous chapter)


| Language      | P-Files| P-LoC  | T-Files| T-LoC  | A-F  |  MLCI | MPI |
| ------------- |-------:| ------:|-------:|-------:|-----:|------:|----:|
| Java          |     30 |   1612 |      4 |    440 | 2052 |   2.4 |  4  |
| Go            |      8 |   1063 |      7 |    508 | 1571 |   1.9 |  2  | 
| Javascript    |      7 |    674 |      4 |    396 | 1070 |   1.3 |  3  |
| Clojure       |      6 |    612 |      4 |    337 |  949 |   1.1 |  1  |
| Python        |      8 |    530 |      5 |    317 |  847 |   1.0 |  1  |


So, Python is the winner in both Marttila Lines of Code Index and Marttila Productivity Index. Clojure doesn't lose that much, productivity being the same, but MLCI being just 10% higher. Javascript loses in productivity quite a bit taking some 3x more implementation time and MLCI is some 30% higher. Go breaks the rule between correlation between MLCI and MPI - MLCI is 90% higher but MPI is only 2x. And Java performs worst of all: MLCI is 140% higher and MPI is 4x.


# Performance

Go performance is excellent since Go compiles to bare metal and there is no runtime between the metal and the language as is the case of other languages (Java & Clojure: compiles to JVM bytecode, and JVM runs the bytecode, Python & Javascript: interpreted languages and the runtime runs the code (though also Python compiles to bytecode).

Go also compiles extremely fast. This was a major design driver when Google designed Go - it had to compile fast since there were major compile time issues with C and C++ because of Google's huge code base (read more in Rob Pike's article ["Go at Google: Language Design in the Service of Software Engineering"](https://talks.golang.org/2012/splash.article), especially chapter 5).


# Concurrency Support

One thing that I did'n have chance to use is the concurrency support provided by the Go language. I read about Go's concurrency support and it seems to be pretty good. The language provides a simple abstraction - goroutines and channels - for concurrency. There is a nice saying among Gophers: "Do not communicate by sharing memory; instead, share memory by communicating." (see Go blog: ["Share Memory By Communicating"](https://blog.golang.org/share-memory-by-communicating)). I.e. in the Java world if you want your threads to share something you share memory and you have to use Java's synchronization mechanisms to provide multi-thread execution. In Go another strategy is used - various entities are collaborating concurrently by sending messages using channels. This is something I definitely want to delve deeper. Clojure also provides pretty good concurrency support since language is immutable by default and you share entities by certain language primitives (like atoms). 


# Go vs Java

Go is so much better that Java in many ways:
- Go compiles to bare metal - Java compiles to byte code which you run in a JVM - overhead.
- Go's error handling is simpler - error handling using specific error entities from functions is more transparent than throwing and catching exceptions (though if done correctly exceptions are also pretty good way to handle errors).
- Go is much more concise than Java - code is easier to handle.
- In Go no class hell - in Java you realize that you are creating class this and class that all the time.
- Go drives you to separate your data (structs) from application logic (functions) - Java drives you to mix both in an unholy mess of classes containing primitive data, methods and references to other classes which contain primitive data, methods and references to other classes... and so on.

But if you need to implement a critical Enterprise system using tens of offshore developers - I still might use Java. Java protects the developers in so many ways when the code base is huge and there are a lot of developers working on the same code base at the same time.


# Conclusions

I fell in love with Go. Go is really a very concise and productive language if you need a robust and performant statically typed language with excellent concurrency support. Much better than Java which compared to Go is verbose, non-productive and concurrency support is far behind Go. I have done quite a lot of C++ and Java programming and I must say that Go's error handling with idiomatic error entity as paired with the actual return value from functions is really great and simple. Go is definitely going to be my choice of statically typed language in my future projects. But still, if I need to create a quick script, e.g. a surrogate script for handling aws cli calls and process return json - I will choose Python. And if I need to process a lot of data - Clojure. But when you need statically typed language and excellent performance with great concurrency support - Go. 

I started my programming career with programming C. Hacking Go is a bit like coming home, except you don't need to be meticulous with memory allocation / deallocation. I think Go gives all goodies from C programming but takes care of the heavy lifting of what's difficult in C. Go code is really simple and elegant - the language provides the exact support for those things that you really need and doesn't add anything extra to the language (like Einstein put it: "Everything should be made as simple as possible, but not simpler").

The feeling was actually quite amazing. I started my Go hacking with practically zero Go knowledge on Monday, and already on Saturday I felt like all pieces of the puzzle just locked in to the right places and creating code was really fluent and easy - and next Monday evening a week later when I started my Go exercise I was already done. Language should be that easy - Go is.
