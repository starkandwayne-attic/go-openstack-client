package flavors

import (
    "fmt"
    "go_openstack_client/apitestharness"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type FlavorsTestSuite struct{
    ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&FlavorsTestSuite{})

func (t *FlavorsTestSuite) SetUpSuite (c *gocheck.C) {
    t.ApiTestHarness = apitestharness.New("nova", false)
}

func (t *FlavorsTestSuite) Test_List (c *gocheck.C) {
    flavors := New(t.ApiTestHarness.ApiConnection)
    fmt.Println(flavors.List())
}
