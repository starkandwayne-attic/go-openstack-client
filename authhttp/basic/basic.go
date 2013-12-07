package basic

import (
    "fmt"
    "errors"
    "net/http"
    "strings"
    "go-openstack-client/authhttp"
)

func name() string {
    return "Basic"
}

type Credentials map[string]string

func (c Credentials) Name() string {
    return name()
}

func (c Credentials) SignRequest(request *http.Request) *http.Request {
    hash := authentication.EncodeStringToBase64(c["username"] + ":" + c["password"])
    request.Header.Add("Authorization", c.Name() + " " + hash)
    return request
}

func (c Credentials) SignClient(client *http.Client) *http.Client {
    return client
}

type Authenticator struct {
    users func() map[string]string
}

func (a Authenticator) Name() string {
    return name()
}

func (a Authenticator) AuthenticateRequest(request *http.Request) (bool, error) {
    if len(request.Header["Authorization"]) == 0 {
      return false, errors.New("Valid Credentials Not Supplied")
    }
    authHeader := strings.Split(request.Header["Authorization"][0]," ")
    decodedAuthString, err := authentication.DecodeBase64String(authHeader[1])
    if err != nil {
      return false, err
    }
    decodedAuth := strings.Split(decodedAuthString,":")
    if strings.HasPrefix(decodedAuthString,":") || len(decodedAuth) < 2 {
      return false, errors.New("Valid Credentials Not Supplied")
    }
    user := decodedAuth[0]
    storedCreds := a.users()[user]
    if storedCreds == "" {
        return false, errors.New("Could Not Find User")
    }
    if storedCreds != authHeader[1] {
        return false, errors.New("Password Mismatch")
    }
    return true, nil
}

func (a Authenticator) ServeUnauthorizedHTTP(w http.ResponseWriter, request *http.Request, err error) {
    w.Header().Set("WWW-Authenticate", "Basic")
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "<html><body><h1>" + err.Error() + "</h1></body></html>")
}
