package none

import (
    "fmt"
    "net/http"
)

func name() string {
    return "None"
}

type Credentials struct {
}

func (c Credentials) Name() string {
    return name()
}

func (c Credentials) SignRequest(request *http.Request) *http.Request {
    request.Header.Add("Authorization", c.Name())
    return request
}

func (c Credentials) SignClient(client *http.Client) *http.Client {
    return client
}

type Authenticator struct {
}

func (a Authenticator) Name() string {
    return name()
}

func (a Authenticator) AuthenticateRequest(request *http.Request) (bool, error) {
    return true, nil
}

func (a Authenticator) ServeUnauthorizedHTTP(w http.ResponseWriter, request *http.Request, err error) {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "<html><body><h1>" + err.Error() + "</h1></body></html>")
}
