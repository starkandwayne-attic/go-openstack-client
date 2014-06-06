package servers

import (
    "fmt"
    "github.com/starkandwayne/go-openstack-client/apitestharness"
    "github.com/starkandwayne/go-openstack-client/nova/images"
    "github.com/starkandwayne/go-openstack-client/nova/flavors"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ServersTestSuite struct{
    ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&ServersTestSuite{})

func (t *ServersTestSuite) SetUpSuite (c *gocheck.C) {
    t.ApiTestHarness = apitestharness.New("compute", false)
}


func (t *ServersTestSuite) xTest_List (c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    fmt.Println(servers.List())
}

func (t *ServersTestSuite) xTest_Get (c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    s, _ := servers.Get("1b2cfce2-c6e3-4368-8c61-d322ddc7412b")
    fmt.Println(s)
    //fmt.Println(s.Addresses["demonet2"][0].Addr)
}

func (t *ServersTestSuite) Test_Create(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    images := images.New(t.ApiTestHarness.ApiConnection)
    flavors := flavors.New(t.ApiTestHarness.ApiConnection)
    options := make(map[string]interface{})
    //privateNet := make(map[string]string)

    options["keyname"] = "bosh"
    options["userdata"] =
`#!/bin/bash
echo "cloud-user    ALL=(ALL)   NOPASSWD: ALL" >> /etc/sudoers`

    //privateNet["uuid"] = "0503030a-9ab9-4807-a7a7-10c06018f3d8"

    //networksList := make([]map[string]string,0)
    //networksList = append(networksList, privateNet)

    //options["networks"] = networksList
    options["security_groups"] = []SecurityGroup{SecurityGroup{"bosh"}}

    serverImage, _ := images.GetByName("centos")
    serverFlavor, _ := flavors.GetByName("m1.small")
    servers.Create("jrbTestServer",serverImage,serverFlavor,options)
}

func (t *ServersTestSuite) xTest_Delete(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    servers.Delete("b998f8cf-5688-4afc-86c6-8d1c457514ab")
}
