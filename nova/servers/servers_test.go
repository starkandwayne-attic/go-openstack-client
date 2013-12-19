package servers

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
    "git.smf.sh/jrbudnack/go_openstack_client/nova/images"
    "git.smf.sh/jrbudnack/go_openstack_client/nova/flavors"
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
    s := servers.Get("187ac9ba-0aef-4eee-a18f-1f20acbaeb4e")
    //fmt.Println(s)
    fmt.Println(s.Addresses["demonet2"][0].Addr)
}

func (t *ServersTestSuite) Test_Create(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    images := images.New(t.ApiTestHarness.ApiConnection)
    flavors := flavors.New(t.ApiTestHarness.ApiConnection)
    options := make(map[string]interface{})
    options["keyname"] = "bosh"
    options["userdata"] =
`#!/bin/bash
echo "cloud-user    ALL=(ALL)   NOPASSWD: ALL" >> /etc/sudoers`
    serverImage, _ := images.GetByName("centos")
    servers.Create("jrbTestServer",serverImage,flavors.List()[1],options)
}

func (t *ServersTestSuite) xTest_Delete(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    servers.Delete("b998f8cf-5688-4afc-86c6-8d1c457514ab")
}
