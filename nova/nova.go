package nova

import (
    "go-openstack-client/apiconnection"
    "go-openstack-client/nova/images"
    "go-openstack-client/nova/flavors"
    "go-openstack-client/nova/servers"
)

type Nova struct {
    ApiConnection apiconnection.ApiConnection
    Images images.Images
    Flavors flavors.Flavors
    Servers servers.Servers
}

func New(adminurl string, username string, password string, tenantname string) {
    n := Nova{}
    n.ApiConnection = apiconnection.New(adminurl,username,password,tenantname)
    n.Servers = servers.New(n.ApiConnection)
    return n
}

//Example:
//n := nova.New("http://10.150.0.10:35357","boshuser","boshpw","bosh")
//n.Servers.List()
//options := make(map[string]interface{})
//n.Servers.Create("jrbTestServer",n.Images.List()[0],n.Flavors.List()[1],options)

//Check Status
//Create Cinder Volume
//Delete/Deprovision

//BONUS POINTS:
//Upload files
//Run remote command
