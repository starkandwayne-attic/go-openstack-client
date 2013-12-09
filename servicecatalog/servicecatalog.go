package servicecatalog

import (
    "fmt"
    //"encoding/json"
    //"io/ioutil"
    //"net/http"
    //"go-openstack-client/util"
)

type ServiceCatalog struct {
    serviceArray []interface{}
}

func New(inServiceArray []interface{}) ServiceCatalog {
    sc := ServiceCatalog{serviceArray: inServiceArray}
    return sc
}

func (sc ServiceCatalog) Show() {
    fmt.Println(sc.serviceArray)
}

//func (ar AuthResponse) parseAuthToken() string {
//    objectParser := util.JsonNode{}
//    objectParser = ar.jsonNode["access"].(map[string]interface{})
//    objectParser = objectParser["token"].(map[string]interface{})
//    return objectParser["id"].(string)
//}
