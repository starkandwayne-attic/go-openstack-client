package cinder

import (
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
    "git.smf.sh/jrbudnack/go_openstack_client/cinder/volumes"
)

type Cinder struct {
    ApiConnection apiconnection.ApiConnection
    Volumes volumes.Volumes
}

func New(adminurl string, username string, password string, tenantname string) Cinder {
    c := Cinder{}
    c.ApiConnection = apiconnection.New(adminurl,"volume",username,password,tenantname)
    c.Volumes = volumes.New(c.ApiConnection)
    return c
}


//Example:

//n := nova.New("http://10.150.0.10:35757","boshuser","boshpw","bosh")
//n.Servers.List()
