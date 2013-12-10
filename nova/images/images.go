package images

import (
    _"fmt"
    "go-openstack-client/apiconnection"
)

type Images struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Images {
    images := Images{apiConnection: apiConnection}
    return images
}

func (s *Images) List() string {
    return s.apiConnection.Get("/images")
}
