package servers

import (
    _"fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
    "go-openstack-client/nova/flavors"
    "go-openstack-client/nova/images"
)

type Servers struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Servers {
    servers := Servers{apiConnection: apiConnection}
    return servers
}

func (s *Servers) List() string {
    return string(s.apiConnection.Get("/servers"))
}

func (s *Servers) Create(name string, image images.Image, flavor flavors.Flavor, options map[string]interface{}) {
    createRequest := make(map[string]interface{})
    serverRequest := make(map[string]interface{})

    serverRequest["name"] = name
    serverRequest["imageRef"] = image.Id
    serverRequest["flavorRef"] = flavor.Id

    createRequest["server"] = serverRequest

    req, _ := json.Marshal(createRequest)
    s.apiConnection.Post("/servers",string(req))
}
