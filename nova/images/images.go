package images

import (
    _"fmt"
    "encoding/json"
    "errors"
    "github.com/starkandwayne/go-openstack-client/apiconnection"
)

type Images struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Images {
    images := Images{apiConnection: apiConnection}
    return images
}

func (s *Images) List() ([]Image, error) {
    type ImagesNode struct {
        Images []Image `json:"images"`
    }
    images := ImagesNode{}
    res, err := s.apiConnection.Get("/images")
    if err != nil {
        return make([]Image,0), err
    }
    json.Unmarshal(res, &images)
    return images.Images, nil
}

func (s *Images) GetByName(name string) (Image, error) {
    images, err := s.List()
    if err != nil {
        return Image{}, err
    }
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
