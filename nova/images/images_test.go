package images

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ImagesTestSuite struct{
    ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&ImagesTestSuite{})

func (t *ImagesTestSuite) SetUpSuite (c *gocheck.C) {
    t.ApiTestHarness = apitestharness.New("compute", false)
}


func (t *ImagesTestSuite) Test_List (c *gocheck.C) {
    images := New(t.ApiTestHarness.ApiConnection)
    fmt.Println(images.List())
}

func (t *ImagesTestSuite) Test_GetByName (c *gocheck.C) {
    images := New(t.ApiTestHarness.ApiConnection)
    fmt.Println(images.GetByName("f17-jeos"))
}
