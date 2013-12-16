package volumes

import (
    "fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
)

type Volumes struct {
    apiConnection apiconnection.ApiConnection
}

type Volume struct {
    Id string
    DisplayName string `json:"display_name"`
    SizeInGB float64 `json:"size"`
    Status string
    Attachments []Attachment
    AvailabilityZone string `json:"availability_zone"`
    Bootable bool
    CreatedAt string `json:"created_at"`
    DisplayDescription string
    VolumeType string `json:"volume_type"`
    SnapshotId string `json:"snapshot_id"`
    SourceVolId string `json:"source_volid"`
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

func (s *Volumes) Get(id string) Volume {
    type VolumeNode struct {
        Volume Volume `json:"volume"`
    }
    volume := VolumeNode{}
    json.Unmarshal(s.apiConnection.Get("/volumes/" + id), &volume)
    return volume.Volume
}

func (vol *Volumes) Create(name string, sizeInGB float64, options map[string]interface{}) Volume {
    createRequest := make(map[string]interface{})
    volumeRequest := make(map[string]interface{})

    volumeRequest["name"] = name
    volumeRequest["size"] = sizeInGB

    createRequest["volume"] = volumeRequest

    type VolumeNode struct {
        Volume Volume `json:"volume"`
    }
    volume := VolumeNode{}

    req, _ := json.Marshal(createRequest)
    json.Unmarshal(vol.apiConnection.Post("/volumes",string(req)), &volume)
    return volume.Volume
}

func (vol *Volumes) Attach(volumeId string, instanceId string, mountPoint string, options map[string]interface{}) {
    action := make(map[string]interface{})
    attachAction := make(map[string]interface{})

    mode := "rw"
    _, hasMode := options["mode"]
    if hasMode {
        mode = options["mode"].(string)
    }

    attachAction["instance_uuid"] = instanceId
    attachAction["mountpoint"] = mountPoint
    attachAction["mode"] = mode

    action["os-attach"] = attachAction

    req, _ := json.Marshal(action)
    fmt.Println(string(vol.apiConnection.Post("/volumes/" + volumeId + "/action",string(req))))
}

func (vol *Volumes) Detach(volumeId string) {
    action := make(map[string]interface{})
    action["os-detach"] = nil

    req, _ := json.Marshal(action)
    fmt.Println(string(vol.apiConnection.Post("/volumes/" + volumeId + "/action",string(req))))
}

func (vol *Volumes) Delete(volumeId string) {
    fmt.Println(string(vol.apiConnection.Delete("/volumes/" + volumeId )))
}
