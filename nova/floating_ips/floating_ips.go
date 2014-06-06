package floating_ips

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"github.com/starkandwayne/go-openstack-client/apiconnection"
	"strconv"
)

type FloatingIps struct {
	apiConnection apiconnection.ApiConnection
}

func New(apiConnection apiconnection.ApiConnection) FloatingIps {
	floatingIps := FloatingIps{apiConnection: apiConnection}
	return floatingIps
}

func (s *FloatingIps) List() ([]FloatingIp, error) {
	type FloatingIpsNode struct {
		FloatingIps []FloatingIp `json:"floating_ips"`
	}
	floatingIps := FloatingIpsNode{}
	res, err := s.apiConnection.Get("/os-floating-ips")
	if err != nil {
		return make([]FloatingIp, 0), err
	}
	json.Unmarshal(res, &floatingIps)
	return floatingIps.FloatingIps, nil
}

func (s *FloatingIps) Create() (FloatingIp, error) {
	type FloatingIpNode struct {
		FloatingIp FloatingIp `json:"floating_ip"`
	}
	floatingIpNode := FloatingIpNode{}
	res, err := s.apiConnection.Post("/os-floating-ips", "")
	if err != nil {
		return FloatingIp{}, err
	}
	json.Unmarshal(res, &floatingIpNode)
	return floatingIpNode.FloatingIp, err
}

func (s *FloatingIps) Delete(id int) error {
	deleteEndpoint := "/os-floating-ips/" + strconv.Itoa(id)
	_, err := s.apiConnection.Delete(deleteEndpoint)
	if err != nil {
		return err
	}
	return nil
}

func (s *FloatingIps) GetById(id int) (FloatingIp, error) {
	floatingIps, err := s.List()
	if err != nil {
		return FloatingIp{}, err
	}
	for _, floatingIp := range floatingIps {
		if floatingIp.Id == id {
			return floatingIp, nil
		}
	}
	return FloatingIp{}, errors.New("Floating Ip not found.")
}

func (s *FloatingIps) CreateAndAttachToServer(serverId string) error {
	floatingIp, err := s.Create()
	if err != nil {
		return err
	}
	err = s.AttachToServer(serverId, floatingIp.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *FloatingIps) AttachToServer(serverId string, floatingIpId int) error {
	serverEndpoint := "/servers/" + serverId + "/action"
	addIpRequest := make(map[string]interface{})
	address := make(map[string]string)
	floatingIp, err := s.GetById(floatingIpId)
	if err != nil {
		return err
	}
	address["address"] = floatingIp.Ip
	addIpRequest["addFloatingIp"] = address
	req, _ := json.Marshal(addIpRequest)
	_, err = s.apiConnection.Post(serverEndpoint, string(req))
	if err != nil {
		return err
	}
	return nil
}

type FloatingIp struct {
	Id         int    `json:"id"`
	InstanceId string `json:"instance_id"`
	Ip         string `json:"ip"`
	Pool       string `json:"pool"`
}
