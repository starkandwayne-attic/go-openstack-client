package volumes

import (
    _"fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
)

type Volumes struct {
    apiConnection apiconnection.ApiConnection
}

type Volume struct {
    Id string
    DisplayName string
    SizeInGB float64
    Status string
    Attachments []Attachment
    AvailabilityZone string
    Bootable bool
    CreatedAt string
    DisplayDescription string
    VolumeType string
    SnapshotId string
    SourceVolId string
    Metadata interface{}
}

type Attachment struct {
    Id string
    VolumeId string
    ServerId string
    Device string
}

func New(apiConnection apiconnection.ApiConnection) Volumes {
    images := Volumes{apiConnection: apiConnection}
    return images
}

func (vol *Volumes) List() []Volume {
    type VolumesNode struct {
        Volumes []Volume `json:"volumes"`
    }
    volumes := VolumesNode{}
    json.Unmarshal(vol.apiConnection.Get("/volumes"), &volumes)
    return volumes.Volumes
}

func (vol *Volumes) Create(name string, sizeInGB float64, options map[string]interface{}) {
    createRequest := make(map[string]interface{})
    volumeRequest := make(map[string]interface{})

    volumeRequest["name"] = name
    volumeRequest["size"] = sizeInGB

    createRequest["volume"] = volumeRequest

    req, _ := json.Marshal(createRequest)
    vol.apiConnection.Post("/volumes",string(req))
}
