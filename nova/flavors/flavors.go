package flavors

import (
    _"fmt"
    "encoding/json"
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
)

type Flavors struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Flavors {
    images := Flavors{apiConnection: apiConnection}
    return images
}

func (s *Flavors) List() []Flavor {
    type FlavorsNode struct {
        Flavors []Flavor `json:"flavors"`
    }
    flavors := FlavorsNode{}
    json.Unmarshal(s.apiConnection.Get("/flavors"), &flavors)
    return flavors.Flavors
}

type Flavor struct {
    Id string
    Name string
    Links []interface{}
}
