package v2

import (
    "bytes"
    "fmt"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "strings"
    "github.com/starkandwayne/go-openstack-client/authhttp"
    "github.com/starkandwayne/go-openstack-client/authresponse"
    "github.com/starkandwayne/go-openstack-client/util"
)

func name() string {
    return "V2"
}

type Credentials map[string]interface{}

func (c Credentials) Name() string {
    return name()
}

func (c Credentials) SignRequest(request *http.Request) *http.Request {
    credentialsMap := util.JsonNode{}
    authMap := util.JsonNode{}
    tokenMap := util.JsonNode{}
    userMap := util.JsonNode{}

    if len(request.Header["X-Auth-Token"]) == 0 {
        _, hasToken := c["token"]
        _, hasServiceCatalog := c["serviceCatalog"]
        _, hasTenantId := c["tenantId"]
        _, hasTenantName := c["tenantName"]
        doPost := false

        //If no Token was provided, we need to log in using a username and
        //password
        if hasToken == false {
            userMap["username"] =  c["username"]
            userMap["password"] = c["password"]
            authMap["passwordCredentials"] = userMap
            doPost = true
        }
        //If the Service Catalog has not been fetched, and we have been provided
        //a Tenant, add it to the auth structure
        if hasServiceCatalog == false &&
            (hasTenantId || hasTenantName) {

            if hasToken {
                //Add token to the credentials map
                tokenMap["id"] = interface{}(c["token"])
                authMap["token"] = tokenMap
            }
            if hasTenantId {
                authMap["tenantId"] = c["tenantId"]
            } else if hasTenantName {
                authMap["tenantName"] = c["tenantName"]
            }
            doPost = true
        }
        //Perform the POST only if there is information we need to fetch.
        if doPost {
            credentialsMap["auth"] = authMap
            requestBody, _ := json.Marshal(credentialsMap)
            buf := ioutil.NopCloser(bytes.NewBufferString(string(requestBody)))
            resp, _ := http.Post("http://" + request.URL.Host + "/v2.0/tokens", "application/json", buf)
            ar := authresponse.New(resp)
            c["token"] = ar.Token
            c["serviceCatalog"] = ar.ServiceCatalog
        }
        //Add the authentication header to the token
        if hasToken {
            request.Header.Add("X-Auth-Token", c["token"].(string))
        }
    }
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
    w.Header().Set("WWW-Authenticate", "V2")
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "<html><body><h1>" + err.Error() + "</h1></body></html>")
}
