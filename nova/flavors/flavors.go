package flavors

import (
    _"fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
)

type Flavors struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Flavors {
    images := Flavors{apiConnection: apiConnection}
    return images
}

func (s *Flavors) List() []Flavor {
    flavors := make(map[string]interface{})
    json.Unmarshal(s.apiConnection.Get("/flavors"), &flavors)
    flavorList := make([]Flavor,0)
    for _, v := range flavors["flavors"].([]interface{}) {
        flavor := v.(map[string]interface{})
        newFlavor := Flavor{Id: flavor["id"].(string), Name: flavor["name"].(string), Links: flavor["links"].([]interface{})}
        flavorList = append(flavorList,newFlavor)
    }
    return flavorList
}

type Flavor struct {
    Id string
    Name string
    Links []interface{}
}
