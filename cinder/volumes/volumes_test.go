package volumes

import (
    "fmt"
    "github.com/starkandwayne/go-openstack-client/apitestharness"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type VolumesTestSuite struct{
    ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&VolumesTestSuite{})

func (t *VolumesTestSuite) SetUpSuite (c *gocheck.C) {
    t.ApiTestHarness = apitestharness.New("volume", false)
}


func (t *VolumesTestSuite) xTest_List (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    volumeList, _ := volumes.List()
    fmt.Println(volumeList)
}

func (t *VolumesTestSuite) xTest_Get (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    v, _ := volumes.Get("56bd7e8b-eee3-4624-a31d-e0f22e3eabf0")
    fmt.Println(v)
}


func (t *VolumesTestSuite) xTest_Create (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    options := make(map[string]interface{})
    v, _ := volumes.Create("jrbNewVolume",float64(20),options)
    fmt.Println(v)
    fmt.Println(v.Id)
}

func (t *VolumesTestSuite) xTest_Detach (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    volumes.Detach("4fb2abe4-7eec-4fc1-a889-20e0d208e102")
}

func (t *VolumesTestSuite) Test_Delete (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    volumes.Delete("4fb2abe4-7eec-4fc1-a889-20e0d208e102")
    volumes.Delete("aad34f71-996e-4c53-b8d0-3c6429a7ad88")
}
