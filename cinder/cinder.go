package cinder

import (
    "go-openstack-client/apiconnection"
    "go-openstack-client/cinder/volumes"
)

type Nova struct {
    ApiConnection apiconnection.ApiConnection
    Volumes volumes.Volumes
}

func New(adminurl string, username string, password string, tenantname string) {
    v := Volumes{}
    v.ApiConnection = apiconnection.New(adminurl,username,password,tenantname)
    v.Volumes = volumes.New(n.ApiConnection)
    return n
}


//Example:

//n := nova.New("http://10.150.0.10:35757","boshuser","boshpw","bosh")
//n.Servers.List()
