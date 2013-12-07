package authenticator

import (
    _"fmt"
    "errors"
    "net/http"
    "strings"
)

// Server Authenticator - Authorize HTTP request from client
type Authenticator interface {
    AuthenticateRequest(request *http.Request) (bool, error)
    ServeUnauthorizedHTTP(w http.ResponseWriter, request *http.Request, err error)
    Name() string
}

type Authenticators map[string]Authenticator

func (a Authenticators) Add(authorizer Authenticator, isDefault bool) {
    a[authorizer.Name()] = authorizer

    if(isDefault) {
        a["default"] = authorizer
    }

    if(a["default"] == nil) {
        a["default"] = authorizer
    }
}

func (a Authenticators) AuthenticateRequest(request *http.Request) (bool, error) {
    authentication := request.Header.Get("Authorization")
    authProvider := "default"
    if authentication != "" {
        authProvider = strings.Split(authentication, " ")[0]
    }
    if a[authProvider] == nil {
        return false, errors.New("Unsupported authentication provider requested.")
    }
    return a[authProvider].AuthenticateRequest(request)
}
