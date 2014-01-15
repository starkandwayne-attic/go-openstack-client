package networks

import (
    _"fmt"
    "errors"
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

func (n *Networks) List() ([]Network, error) {
    type NetworksNode struct {
        Networks []Network `json:"networks"`
    }
    networks := NetworksNode{}
    res, err := n.apiConnection.Get("/v2.0/networks")
    if err != nil {
        return make([]Network,0), err
    }
    json.Unmarshal(res, &networks)
    return networks.Networks, nil
}

func (n *Networks) GetByName(name string) (Network, error) {
    networks, err := n.List()
    if err != nil {
        return Network{}, err
    }
    for _, network := range networks {
        if network.Name == name {
            return network, nil
        }
    }
    return Network{}, errors.New("Network not found.")
}

//func (n *Networks) Get(id string) Network {
    //type VolumeNode struct {
    //    Volume Volume `json:"volume"`
    //}
    //volume := VolumeNode{}
    //json.Unmarshal(s.apiConnection.Get("/volumes/" + id), &volume)
    //return volume.Volume
//}
