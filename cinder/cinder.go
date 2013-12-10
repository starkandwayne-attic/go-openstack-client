package nova

import (
    "go-openstack-client/nova/apiconnection"
    "go-openstack-client/nova/servers"
)

type Nova struct {
    ApiConnection apiconnection.ApiConnection
    Servers servers.Servers
}

func New(adminurl string, username string, password string, tenantname string) {
    n := Nova{}
    n.ApiConnection = apiconnection.New(adminurl,username,password,tenantname)
    n.Servers = servers.New(n.ApiConnection)
    return n
}


//Example:

//n := nova.New("http://10.150.0.10:35757","boshuser","boshpw","bosh")
//n.Servers.List()
