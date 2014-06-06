package quantum

import (
    "github.com/starkandwayne/go-openstack-client/apiconnection"
    "github.com/starkandwayne/go-openstack-client/quantum/networks"
)

type Quantum struct {
    ApiConnection apiconnection.ApiConnection
    Networks networks.Networks
}

func New(adminurl string, username string, password string, tenantname string) Quantum {
    q := Quantum{}
    q.ApiConnection = apiconnection.New(adminurl,"network",username,password,tenantname)
    q.Networks = networks.New(q.ApiConnection)
    return q
}
