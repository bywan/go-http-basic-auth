package basic

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"
)

type BasicAuthenticator struct {
	Validate func(username, password string) bool
	Realm    string
}

func NewAuthenticator(validator func(username, password string) bool, realm string) *BasicAuthenticator {
	return &BasicAuthenticator{validator, realm}
}

func (a *BasicAuthenticator) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.auth(r) {

			next.ServeHTTP(w, r)

		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+a.Realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
		}
	})
}

func (a *BasicAuthenticator) auth(req *http.Request) bool {
	pb := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
	if len(pb) != 2 {
		err := errors.New("Bad auth")
		log.Println("Error", err)
		return false
	}

	if len(pb[1]) == 0 {
		err := errors.New("Authentication required")
		log.Println("Error", err)
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(pb[1])
	if err != nil {
		log.Println("Error bad auth string")
		return false
	}

	lp := strings.SplitN(string(decoded), ":", 2)
	if len(pb) != 2 {
		err := errors.New("Bad auth")
		log.Println("Error", err)
		return false
	}

	return a.Validate(lp[0], lp[1])
}
