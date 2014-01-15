package flavors

import (
    _"fmt"
    "encoding/json"
    "errors"
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
)

type Flavors struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Flavors {
    images := Flavors{apiConnection: apiConnection}
    return images
}

func (s *Flavors) List() ([]Flavor, error) {
    type FlavorsNode struct {
        Flavors []Flavor `json:"flavors"`
    }
    flavors := FlavorsNode{}
    res, err := s.apiConnection.Get("/flavors")

    if err != nil {
        return make([]Flavor,0), err
    }

    json.Unmarshal(res, &flavors)
    return flavors.Flavors, nil
}

func (s *Flavors) GetByName(name string) (Flavor, error) {
    flavors, err := s.List()
    if err != nil {
        return Flavor{}, err
    }

    for _, flavor := range flavors {
        if flavor.Name == name {
            return flavor, nil
        }
    }
    return Flavor{}, errors.New("Flavor not found.")
}

type Flavor struct {
    Id string
    Name string
    Links []interface{}
}
