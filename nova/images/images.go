package images

import (
    _"fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
)

type Images struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Images {
    images := Images{apiConnection: apiConnection}
    return images
}

func (s *Images) List() []Image {
    images := make(map[string]interface{})
    json.Unmarshal(s.apiConnection.Get("/images"), &images)
    imageList := make([]Image,0)
    for _, v := range images["images"].([]interface{}) {
        image := v.(map[string]interface{})
        newImage := Image{Id: image["id"].(string),
                          Name: image["name"].(string),
                          Links: image["links"].([]interface{})}
        imageList = append(imageList,newImage)
    }
    return imageList
}

type Image struct {
    Id string
    Name string
    Links []interface{}
}
