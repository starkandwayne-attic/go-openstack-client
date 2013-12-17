package images

import (
    _"fmt"
    "encoding/json"
    "go_openstack_client/apiconnection"
)

type Images struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Images {
    images := Images{apiConnection: apiConnection}
    return images
}

func (s *Images) List() []Image {
    type ImagesNode struct {
        Images []Image `json:"images"`
    }
    images := ImagesNode{}
    json.Unmarshal(s.apiConnection.Get("/images"), &images)
    return images.Images
}

type Image struct {
    Id string
    Name string
    Links []interface{}
}
