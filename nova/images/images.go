package images

import (
    _"fmt"
    "encoding/json"
    "errors"
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
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

func (s *Images) GetByName(name string) (Image, error) {
    images := s.List()
    for _, image := range images {
        if image.Name == name {
            return image, nil
        }
    }
    return Image{}, errors.New("Image not found.")
}


type Image struct {
    Id string
    Name string
    Links []interface{}
}
