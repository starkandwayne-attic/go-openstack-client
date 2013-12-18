package networks

import (
    _"fmt"
    "encoding/json"
    "git.smf.sh/jrbudnack/go_openstack_client/apiconnection"
)

type Networks struct {
    apiConnection apiconnection.ApiConnection
}

type Network struct {
    Id string
    Status string
    Subnets []string
    Name string
    RouterExternal bool `json:"router:external"`
    TenantId string `json:"tenant_id"`
    AdminStateUp bool `json:"admin_state_up"`
    Shared bool
}

func New(apiConnection apiconnection.ApiConnection) Networks {
    networks := Networks{apiConnection: apiConnection}
    return networks
}

func (n *Networks) List() []Network {
    type NetworksNode struct {
        Networks []Network `json:"networks"`
    }
    networks := NetworksNode{}
    json.Unmarshal(n.apiConnection.Get("/v2.0/networks"), &networks)
    return networks.Networks
}

//func (n *Networks) Get(id string) Network {
    //type VolumeNode struct {
    //    Volume Volume `json:"volume"`
    //}
    //volume := VolumeNode{}
    //json.Unmarshal(s.apiConnection.Get("/volumes/" + id), &volume)
    //return volume.Volume
//}
