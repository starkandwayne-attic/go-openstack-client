package quantum

import (
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
    "git.smf.sh/jrbudnack/go_openstack_client/quantum/networks"
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
