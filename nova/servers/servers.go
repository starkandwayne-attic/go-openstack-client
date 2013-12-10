package servers

import (
    _"fmt"
    "go-openstack-client/nova/apiconnection"
)

type Servers struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Servers {
    servers := Servers{apiConnection: apiConnection}
    return servers
}

func (s *Servers) GetServerList() string {
    return s.apiConnection.Get("/servers")
}