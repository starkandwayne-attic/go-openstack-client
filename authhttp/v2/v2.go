package v2

import (
    "bytes"
    "fmt"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "strings"
    "go-openstack-client/authhttp"
)

func name() string {
    return "V2"
}

type Credentials map[string]string

func (c Credentials) Name() string {
    return name()
}

func (c Credentials) SignRequest(request *http.Request) *http.Request {
    credentialsMap := make(map[string]interface {})
    authMap := make(map[string]interface{})
    tokenMap := make(map[string]interface{})
    userMap := make(map[string]interface{})

    if len(request.Header["X-Auth-Token"]) == 0 {
        _, hasToken := c["token"]
        if hasToken {
            //Add token to the credentials map
            tokenMap["id"] = interface{}(c["token"])
            authMap["token"] = tokenMap
            //Add authentication token header
            request.Header.Add("X-Auth-Token", c["token"])
        } else {
            //Add username and password to the body
            userMap["username"] =  c["username"]
            userMap["password"] = c["password"]
            authMap["passwordCredentials"] = userMap
        }
        _, hasTenantId := c["tenantId"]
        _, hasTenantName := c["tenantName"]
        if hasTenantId {
            authMap["tenantId"] = c["tenantId"]
        } else if hasTenantName {
            authMap["tenantName"] = c["tenantName"]
        }
        credentialsMap["auth"] = authMap
    }
    body, _ := json.Marshal(credentialsMap)
    request.Body = ioutil.NopCloser(bytes.NewBufferString(string(body)))
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
