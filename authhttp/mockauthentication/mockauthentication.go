package mockauthentication

import (
    "fmt"
    "errors"
    "net/http"
    "strings"
)

func name() string {
    return "TestAUTH"
}

type Credentials map[string]string

func (c Credentials) Name() string {
    return name()
}

func (c Credentials) SignRequest(request *http.Request) *http.Request {
    var authString string = c.Name() + " " + c["user"] + ":" + c["password"]
    request.Header.Add("Authorization", authString)
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
    authentication := request.Header["Authorization"]
    if len(authentication) == 0 {
      return false, errors.New("Unauthorized")
    }
    if strings.Join(authentication," ") == "TestAUTH Jeremy:PicklesRGR8" {
      return true, nil
    }
    return false, errors.New("Unauthorized")
}

func (a Authenticator) ServeUnauthorizedHTTP(w http.ResponseWriter, request *http.Request, err error) {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "<html><body><h1>" + err.Error() + "</h1></body></html>")
}
