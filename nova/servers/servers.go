package servers

import (
    "fmt"
    "encoding/json"
    "go-openstack-client/apiconnection"
    "go-openstack-client/nova/flavors"
    "go-openstack-client/nova/images"
    "go-openstack-client/cinder/volumes"
)

type Servers struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Servers {
    servers := Servers{apiConnection: apiConnection}
    return servers
}

func (s *Servers) List() []Server {
    type ServersNode struct {
        Servers []Server `json:"servers"`
    }
    servers := ServersNode{}
    json.Unmarshal(s.apiConnection.Get("/servers/detail"), &servers)
    return servers.Servers
}

func (s *Servers) Get(id string) Server {
    type ServerNode struct {
        Server Server `json:"server"`
    }
    server := ServerNode{}
    json.Unmarshal(s.apiConnection.Get("/servers/" + id), &server)
    return server.Server
}

func (s *Servers) Create(name string, image images.Image, flavor flavors.Flavor, options map[string]interface{}) Server {
    createRequest := make(map[string]interface{})
    serverRequest := make(map[string]interface{})

    serverRequest["name"] = name
    serverRequest["imageRef"] = image.Id
    serverRequest["flavorRef"] = flavor.Id

    createRequest["server"] = serverRequest

    type ServerNode struct {
        Server Server `json:"server"`
    }
    server := ServerNode{}
    req, _ := json.Marshal(createRequest)
    json.Unmarshal(s.apiConnection.Post("/servers",string(req)),&server)
    return server.Server
}

type Server struct {
    Id string
    Name string
    Created string
    Status string
    Updated string
    HostId string
    Addresses map[string][]Address
    Links []Link
    KeyName string `json:"key_name"`
    Image images.Image
    TaskState string `json:"OS-EXT-STS:task_state"`
    VMState string `json:"OS-EXT-STS:vm_state"`
    Flavor flavors.Flavor
    SecurityGroups []SecurityGroup `json:"security_groups"`
    AvailabilityZone string `json:"OS-EXT-AZ:availability_zone"`
    UserID string `json:"user_id"`
    TenantID string `json:"tenant_id"`
    DiskConfig string `json:"OS-DCF:diskConfig"`
    AccessIPV4 string
    AccessIPV6 string
    Progress int
    PowerState int `json:"OS-EXT-STS:power_state"`
    ConfigDrive string `json:"config_drive"`
    Metadata map[string]string
}

type Address struct {
    Version int
    Addr string
    Type string `json:"OS-EXT-IPS:type"`
}

type Link struct {
    Href string
    Rel string
}

type SecurityGroup struct {
    Name string
}
