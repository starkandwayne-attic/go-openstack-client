package apiconnection

import (
    _"fmt"
    "io/ioutil"
    "go-openstack-client/authhttp/client"
    "go-openstack-client/authhttp/v2"
    "go-openstack-client/servicecatalog"
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
    endpointQuery["name"] = ac.ServiceName

    ac.authHttpClient = client.New(adminCreds, ac.AdminUrl)
    ac.authHttpClient.Get("/")

    sc := adminCreds["serviceCatalog"].(servicecatalog.ServiceCatalog)
    ac.ApiConnectionUrl = sc.GetEndpoint(endpointQuery)

    ac.novaHttpClient = client.New(adminCreds, ac.ApiConnectionUrl)
}

func (ac *ApiConnection) Get(endpointURL string) string {
    res, _ := ac.novaHttpClient.Get(endpointURL)
    body, _ := ioutil.ReadAll(res.Body)
    return string(body)
}
