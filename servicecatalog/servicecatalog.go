package servicecatalog

import (
    _"fmt"
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

func (sc *ServiceCatalog) GetEndpoint(queryParameters map[string]string) string {
    //Query Parameters:
    //  urltype (admin/internal/public)
    //  name (glance/nova/...)
    //  region
    //  type (image/compute/...)
    urlType, urlTypeProvided := queryParameters["urltype"]
    if !urlTypeProvided {
        return ""
    }
    urlType = urlType + "URL"
    for _, svc := range sc.serviceArray {
        svcMap := svc.(map[string]interface{})
        endpoints := svcMap["endpoints"].([]interface{})
        for _, edp := range endpoints {
            meetsAll := true
            edpMap := edp.(map[string]interface{})
            for k,v := range queryParameters {
                _, keyExists := edpMap[k]
                if keyExists && edpMap[k] != v {
                    meetsAll = false
                    break
                }
            }
            if meetsAll {
                return edpMap[urlType].(string)
            }
        }
    }
    return ""
}
