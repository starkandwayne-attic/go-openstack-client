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
    s := servers.Get("c4c1630b-de71-44b2-aea4-e52067e149fb")
    fmt.Println(s)
}

func (t *ServersTestSuite) Test_Create(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    images := images.New(t.ApiTestHarness.ApiConnection)
    flavors := flavors.New(t.ApiTestHarness.ApiConnection)
    options := make(map[string]interface{})
    options["keyname"] = "bosh"

    servers.Create("jrbTestServer",images.List()[0],flavors.List()[1],options)
}

func (t *ServersTestSuite) xTest_Delete(c *gocheck.C) {
    servers := New(t.ApiTestHarness.ApiConnection)
    servers.Delete("b998f8cf-5688-4afc-86c6-8d1c457514ab")
}
