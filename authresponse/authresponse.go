package authresponse

import (
    _"fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "git.smf.sh/jrbudnack/go_openstack_client/servicecatalog"
    "git.smf.sh/jrbudnack/go_openstack_client/util"
)

type AuthResponse struct {
    jsonNode util.JsonNode
    Token string
    ServiceCatalog servicecatalog.ServiceCatalog
}

func New(inResponse *http.Response) AuthResponse {
    responseBody, _ := ioutil.ReadAll(inResponse.Body)
    authResponse := util.JsonNode{}
    _ = json.Unmarshal(responseBody,&authResponse)

    ar := AuthResponse{jsonNode: authResponse}
    ar.parseAuthToken()
    ar.parseServiceCatalog()
    return ar
}

func (ar *AuthResponse) parseAuthToken() {
    objectParser := util.JsonNode{}
    objectParser = ar.jsonNode["access"].(map[string]interface{})

    _, hasToken := objectParser["token"]
    if hasToken {
        objectParser = objectParser["token"].(map[string]interface{})
        ar.Token = objectParser["id"].(string)
    }
}

func (ar *AuthResponse) parseServiceCatalog() {
    objectParser := util.JsonNode{}
    objectParser = ar.jsonNode["access"].(map[string]interface{})

    _, hasSC := objectParser["serviceCatalog"]
    if hasSC {
        ar.ServiceCatalog = servicecatalog.New(objectParser["serviceCatalog"].([]interface{}))
    }
}
