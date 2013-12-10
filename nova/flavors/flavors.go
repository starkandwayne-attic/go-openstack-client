package flavors

import (
    _"fmt"
    "go-openstack-client/apiconnection"
)

type Flavors struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Flavors {
    images := Flavors{apiConnection: apiConnection}
    return images
}

func (s *Flavors) List() string {
    return s.apiConnection.Get("/flavors")
}
