# Simple HTTP Basic Authentication in Go

[![Travis](https://img.shields.io/travis/bywan/go-http-basic-auth.svg?style=flat-square)](https://travis-ci.org/bywan/go-http-basic-auth)

This is an implementation of HTTP Basic authentication in Go.

## Usage

```go

package main

import (
	"log"
	"net/http"

	auth "github.com/bywan/go-http-basic-auth"
)

func myValidator(username, password string) bool {
	// Here is your validator logic
	return username == "foo" && password == "bar"
}

func main() {
	authenticator := auth.NewAuthenticator(myValidator, "myRealm")
	appHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	http.Handle("/", authenticator.Wrap(appHandler))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

```