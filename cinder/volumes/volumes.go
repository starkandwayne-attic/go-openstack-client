package volumes

import (
    _"fmt"
    "encoding/json"
    "github.com/starkandwayne/go-openstack-client/apiconnection"
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

func (vol *Volumes) List() ([]Volume, error) {
    type VolumesNode struct {
        Volumes []Volume `json:"volumes"`
    }
    volumes := VolumesNode{}
    res, err := vol.apiConnection.Get("/volumes")
    if err != nil {
        return make([]Volume,0), err
    }

    json.Unmarshal(res, &volumes)
    return volumes.Volumes, nil
}

func (vol *Volumes) Get(id string) (Volume, error) {
    type VolumeNode struct {
        Volume Volume `json:"volume"`
    }
    volume := VolumeNode{}
    res, err := vol.apiConnection.Get("/volumes/" + id)
    if err != nil {
        return Volume{}, err
    }
    json.Unmarshal(res, &volume)
    return volume.Volume, nil
}

func (vol *Volumes) Create(name string, sizeInGB float64, options map[string]interface{}) (Volume, error) {
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
    res, err := vol.apiConnection.Post("/volumes", string(req))
    if err != nil {
        return Volume{}, err
    }
    json.Unmarshal(res, &volume)
    return volume.Volume, nil
}

func (vol *Volumes) Attach(volumeId string, instanceId string, mountPoint string, options map[string]interface{}) error {
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
    _, err := vol.apiConnection.Post("/volumes/" + volumeId + "/action", string(req))

    return err
}

func (vol *Volumes) Detach(volumeId string) error {
    action := make(map[string]interface{})
    action["os-detach"] = nil

    req, _ := json.Marshal(action)
    _, err := vol.apiConnection.Post("/volumes/" + volumeId + "/action", string(req))
    return err
}

func (vol *Volumes) Delete(volumeId string) error {
    _, err := vol.apiConnection.Delete("/volumes/" + volumeId)
    return err
}
