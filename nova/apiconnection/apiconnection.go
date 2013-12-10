package apiconnection

import (
    "fmt"
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
}

func New(adminurl string, username string, password string, tenantname string) ApiConnection {
    ac := ApiConnection{AdminUrl: adminurl, Username: username, Password: password, TenantName: tenantname}
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
    endpointQuery["servicename"] = "nova"

    ac.authHttpClient = client.New(adminCreds, ac.AdminUrl)
    ac.authHttpClient.Get("/")

    sc := adminCreds["serviceCatalog"].(servicecatalog.ServiceCatalog)
    ac.ApiConnectionUrl = sc.GetEndpoint(endpointQuery)

    ac.novaHttpClient = client.New(adminCreds, ac.ApiConnectionUrl)
    res, _ := ac.novaHttpClient.Get("/servers")
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))
}

func (ac *ApiConnection) PrintServerList() {
    res, _ := ac.novaHttpClient.Get("/servers")
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))
}
