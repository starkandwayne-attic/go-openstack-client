package handler

import (
  _"fmt"
  "github.com/starkandwayne/go-openstack-client/authhttp/authenticator"
  "net/http"
)

type AuthHandler struct {
    authenticators authenticator.Authenticators
    handler http.Handler
}

func New(authenticators authenticator.Authenticators, handler http.Handler) AuthHandler {
    return AuthHandler{authenticators: authenticators, handler: handler}
}

func (a AuthHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
    authorized, err := a.authenticators.AuthenticateRequest(request)
    if authorized {
        a.handler.ServeHTTP(w,request)
    } else {
        a.authenticators["default"].ServeUnauthorizedHTTP(w,request,err)
    }
}


