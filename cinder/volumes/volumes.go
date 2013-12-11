package volumes

import (
    _"fmt"
    "encoding/json"
    "strconv"
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
    volumes := make(map[string]interface{})
    json.Unmarshal(vol.apiConnection.Get("/volumes"), &volumes)
    volumeList := make([]Volume,0)
    for _, v := range volumes["volumes"].([]interface{}) {
        volume := v.(map[string]interface{})
        attachmentList := make([]Attachment,0)
        _, hasAttachments := volume["attachments"]
        if hasAttachments {
            attachments := volume["attachments"].([]interface{})
            for _, a := range attachments {
                attachment := a.(map[string]interface{})
                newAttachment := Attachment{Id: attachment["id"].(string),
                                            VolumeId: attachment["volume_id"].(string),
                                            ServerId: attachment["server_id"].(string),
                                            Device: attachment["device"].(string),
                }
                attachmentList = append(attachmentList,newAttachment)
            }
        }
        bootable, _ := strconv.ParseBool(volume["bootable"].(string))
        newVolume := Volume{Id: volume["id"].(string),
                            DisplayName: volume["display_name"].(string),
                            SizeInGB: volume["size"].(float64),
                            Status: volume["status"].(string),
                            Attachments: attachmentList,
                            AvailabilityZone: volume["availability_zone"].(string),
                            Bootable: bootable,
                            CreatedAt: volume["created_at"].(string),
                            DisplayDescription: volume["display_description"].(string),
                            //VolumeType: volume["volume_type"].(string),
                            //SnapshotId: volume["snapshot_id"].(string),
                            //SourceVolId: volume["source_volid"].(string),
                            Metadata: volume["metadata"].(interface{}),
        }
        volumeList = append(volumeList,newVolume)
    }
    return volumeList
}

//func (vol *Volumes) Create(name string, image images.Image, flavor flavors.Flavor, options map[string]interface{}) {
    //createRequest := make(map[string]interface{})
    //serverRequest := make(map[string]interface{})

    //serverRequest["name"] = name
    //serverRequest["imageRef"] = image.Id
    //serverRequest["flavorRef"] = flavor.Id

    //createRequest["server"] = serverRequest

    //req, _ := json.Marshal(createRequest)
    //s.apiConnection.Post("/servers",string(req))
//}
