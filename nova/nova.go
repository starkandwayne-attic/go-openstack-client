package nova

import (
    "fmt"
    "io/ioutil"
    "go-openstack-client/authhttp/client"
    "go-openstack-client/authhttp/v2"
    "go-openstack-client/servicecatalog"
)

type Nova struct {
    AdminUrl string
    Username string
    Password string
    TenantName string
    authHttpClient client.Client

    NovaUrl string
    novaHttpClient client.Client
}

func New(adminurl string, username string, password string, tenantname string) Nova {
    n := Nova{AdminUrl: adminurl, Username: username, Password: password, TenantName: tenantname}
    n.Connect()
    return n
}

func (n *Nova) Connect() {
    adminCreds := v2.Credentials{}
    adminCreds["username"] = n.Username
    adminCreds["password"] = n.Password
    adminCreds["tenantName"] = n.TenantName

    endpointQuery := make(map[string]string)
    endpointQuery["urltype"] = "public"
    endpointQuery["servicename"] = "nova"

    n.authHttpClient = client.New(adminCreds, n.AdminUrl)
    n.authHttpClient.Get("/")

    sc := adminCreds["serviceCatalog"].(servicecatalog.ServiceCatalog)
    n.NovaUrl = sc.GetEndpoint(endpointQuery)

    n.novaHttpClient = client.New(adminCreds, n.NovaUrl)
    res, _ := n.novaHttpClient.Get("/servers")
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))
}

func (n *Nova) PrintServerList() {
    res, _ := n.novaHttpClient.Get("/servers")
    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))
}
