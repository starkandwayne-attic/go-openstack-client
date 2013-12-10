package volumes

import (
    _"fmt"
    "go-openstack-client/apiconnection"
)

type Volumes struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Volumes {
    images := Volumes{apiConnection: apiConnection}
    return images
}

func (s *Volumes) List() string {
    return s.apiConnection.Get("/volumes")
}
