package apiconnection

import (
    _"fmt"
    "io/ioutil"
    "github.com/starkandwayne/go-openstack-client/authhttp/client"
    "github.com/starkandwayne/go-openstack-client/authhttp/v2"
    "github.com/starkandwayne/go-openstack-client/servicecatalog"
)

type ApiConnection struct {
    AdminUrl string
    Username string
    Password string
    TenantName string
    authHttpClient client.Client

    ApiConnectionUrl string
    novaHttpClient client.Client
    ServiceName string
}

func New(adminurl string, servicename string, username string, password string, tenantname string) ApiConnection {
    ac := ApiConnection{AdminUrl: adminurl, ServiceName: servicename, Username: username, Password: password, TenantName: tenantname}
    ac.Connect()
    return ac
}

func (ac *ApiConnection) Connect() {
    adminCreds := v2.Credentials{}
    adminCreds["username"] = ac.Username
    adminCreds["password"] = ac.Password
    adminCreds["tenantName"] = ac.TenantName

    endpointQuery := make(map[string]string)
    endpointQuery["urltype"] = "public"
    endpointQuery["type"] = ac.ServiceName

    ac.authHttpClient = client.New(adminCreds, ac.AdminUrl)
    ac.authHttpClient.Get("/")

    sc := adminCreds["serviceCatalog"].(servicecatalog.ServiceCatalog)
    //fmt.Println(sc)
    ac.ApiConnectionUrl = sc.GetEndpoint(endpointQuery)

    ac.novaHttpClient = client.New(adminCreds, ac.ApiConnectionUrl)
}

func (ac *ApiConnection) Get(endpointURL string) ([]byte, error) {
    res, httpErr := ac.novaHttpClient.Get(endpointURL)

    if httpErr != nil {
        return make([]byte,0), httpErr
    }

    resBody, readErr := ioutil.ReadAll(res.Body)

    if readErr != nil {
        return make([]byte,0), readErr
    }

    return resBody, nil
}

func (ac *ApiConnection) Post(endpointURL string, body string) ([]byte, error) {
    res, httpErr := ac.novaHttpClient.Post(endpointURL,body)

    if httpErr != nil {
        return make([]byte,0), httpErr
    }

    resBody, readErr := ioutil.ReadAll(res.Body)

    if readErr != nil {
        return make([]byte,0), readErr
    }

    return resBody, nil
}

func (ac *ApiConnection) Delete(endpointURL string) ([]byte, error) {
    res, httpErr := ac.novaHttpClient.Delete(endpointURL)

    if httpErr != nil {
        return make([]byte,0), httpErr
    }

    resBody, readErr := ioutil.ReadAll(res.Body)

    if readErr != nil {
        return make([]byte,0), readErr
    }

    return resBody, nil
}
