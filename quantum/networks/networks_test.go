package networks

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type NetworksTestSuite struct{
    ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&NetworksTestSuite{})

func (t *NetworksTestSuite) SetUpSuite (c *gocheck.C) {
    t.ApiTestHarness = apitestharness.New("network", false)
}


func (t *NetworksTestSuite) Test_List (c *gocheck.C) {
    networks := New(t.ApiTestHarness.ApiConnection)
    networkList := networks.List()
    fmt.Println(networkList)
}

//func (t *NetworksTestSuite) xTest_Get (c *gocheck.C) {
//    networks := New(t.ApiTestHarness.ApiConnection)
//    v := networks.Get("56bd7e8b-eee3-4624-a31d-e0f22e3eabf0")
//    fmt.Println(v)
//}
