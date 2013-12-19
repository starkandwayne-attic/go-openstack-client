package volumes

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
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
    volumeList := volumes.List()
    fmt.Println(volumeList)
}

func (t *VolumesTestSuite) xTest_Get (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    v := volumes.Get("56bd7e8b-eee3-4624-a31d-e0f22e3eabf0")
    fmt.Println(v)
}


func (t *VolumesTestSuite) xTest_Create (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    options := make(map[string]interface{})
    v := volumes.Create("jrbNewVolume",float64(20),options)
    fmt.Println(v)
    fmt.Println(v.Id)
}

func (t *VolumesTestSuite) Test_Detach (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    volumes.Detach("303152e4-14f7-47c0-8485-8efd421f055d")
}

func (t *VolumesTestSuite) xTest_Delete (c *gocheck.C) {
    volumes := New(t.ApiTestHarness.ApiConnection)
    volumes.Delete("34e1763d-47a6-4802-b47f-b9f2be956a2f")
}
