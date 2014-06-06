package servers

import (
    "fmt"
    "encoding/base64"
    "encoding/json"
    "github.com/starkandwayne/go-openstack-client/apiconnection"
    "github.com/starkandwayne/go-openstack-client/nova/flavors"
    "github.com/starkandwayne/go-openstack-client/nova/images"
)

type Servers struct {
    apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) Servers {
    servers := Servers{apiConnection: apiConnection}
    return servers
}

func (s *Servers) List() ([]Server, error) {
    type ServersNode struct {
        Servers []Server `json:"servers"`
    }
    servers := ServersNode{}
    res, err := s.apiConnection.Get("/servers/detail")
    if err != nil {
        return make([]Server,0), err
    }
    json.Unmarshal(res, &servers)
    return servers.Servers, nil
}

func (s *Servers) Get(id string) (Server, error) {
    type ServerNode struct {
        Server Server `json:"server"`
    }
    server := ServerNode{}
    res, err := s.apiConnection.Get("/servers/" + id)
    if err != nil {
        return Server{}, err
    }
    json.Unmarshal(res, &server)
    return server.Server, nil
}

func (s *Servers) Create(name string, image images.Image, flavor flavors.Flavor, options map[string]interface{}) (Server, error) {
    createRequest := make(map[string]interface{})
    serverRequest := make(map[string]interface{})

    serverRequest["name"] = name
    serverRequest["imageRef"] = image.Id
    serverRequest["flavorRef"] = flavor.Id

    _, hasKeyName := options["keyname"]

    if hasKeyName {
        serverRequest["key_name"] = options["keyname"]
    }

    _, hasUserData := options["userdata"]

    if hasUserData {
        encodedUserData := base64.StdEncoding.EncodeToString([]byte(options["userdata"].(string)))
        serverRequest["user_data"] = encodedUserData
    }

    _, hasSecurityGroups := options["security_groups"]

    if hasSecurityGroups{
        serverRequest["security_groups"] = options["security_groups"]
    }

    _, hasNetworks := options["networks"]

    if hasNetworks {
        serverRequest["networks"] = options["networks"]
    }

    createRequest["server"] = serverRequest

    type ServerNode struct {
        Server Server `json:"server"`
    }
    server := ServerNode{}
    req, _ := json.Marshal(createRequest)
    res, err := s.apiConnection.Post("/servers", string(req))
    if err != nil {
        fmt.Println(err)
        return Server{}, err
    }

    json.Unmarshal(res, &server)
    return server.Server, nil
}

func (s *Servers) Delete(id string) (error) {
    _, err := s.apiConnection.Delete("/servers/" + id)
    return err
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
