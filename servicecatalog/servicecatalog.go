package servicecatalog

import (
    _"fmt"
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
    for i := range sc.serviceArray {
        svcMap := sc.serviceArray[i].(map[string]interface{})
        endpoints := svcMap["endpoints"].([]interface{})
        meetsAll := true
        for k, v := range queryParameters {
            _, keyExists := svcMap[k]
            if keyExists && svcMap[k] != v {
                meetsAll = false
            }
        }
        if meetsAll {
            for _, edp := range endpoints {
                edpMap := edp.(map[string]interface{})
                return edpMap[urlType].(string)
            }
        }
    }
    return ""
}
