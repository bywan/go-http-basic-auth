# Simple HTTP Basic Authentication in Go

[![Travis](https://img.shields.io/travis/bywan/go-http-basic-auth.svg?style=flat-square)](https://travis-ci.org/bywan/go-http-basic-auth)

This is an implementation of HTTP Basic authentication in Go.

## Usage

```go

package main

import (
    auth "github.com/bywan/go-http-auth"
    "log"
    "net/http"
)

func myValidator(username, password string) bool {
    // Here is your validator logic
    return username == "foo" && password == "bar"
}

func main() {
	auth := NewAuthenticator(myValidator, "myRealm")
	appHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    w.WriteHeader(204)
	})

	httptest.NewServer(auth.Wrap(appHandler))

	http.Handle("/", )
	log.Fatal(http.ListenAndServe(":3000", nil))
}

```