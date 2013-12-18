package nova

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
    "git.smf.sh/jrbudnack/go_openstack_client/cinder"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ServersTestSuite struct{
    NovaApiTestHarness apitestharness.ApiTestHarness
    CinderApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&ServersTestSuite{})

func (t *ServersTestSuite) SetUpSuite (c *gocheck.C) {
    t.NovaApiTestHarness = apitestharness.New("compute", false)
    t.CinderApiTestHarness = apitestharness.New("volume", false)
}

func (t *ServersTestSuite) Test_CreateAFreakingServer (c *gocheck.C) {
    fmt.Println("Creating volume...")
    n := New(t.NovaApiTestHarness.Url,
             t.NovaApiTestHarness.Username,
             t.NovaApiTestHarness.Password,
             t.NovaApiTestHarness.Tenant)
    cdr := cinder.New(t.CinderApiTestHarness.Url,
                      t.CinderApiTestHarness.Username,
                      t.CinderApiTestHarness.Password,
                      t.CinderApiTestHarness.Tenant)
    fmt.Println("Volume has been created.  Creating server...")
    volumeOptions := make(map[string]interface{})
    v := cdr.Volumes.Create("jrbNewVolume",float64(20),volumeOptions)
    for v.Status != "available" && v.Status != "error" {
        v = cdr.Volumes.Get(v.Id)
    }
    serverOptions := make(map[string]interface{})
    serverOptions["keyname"] = "bosh"
    sourceImage, _ := n.Images.GetByName("BOSH-a7430ae2-e573-412b-8c5d-6d1571cb064a")
    s := n.Servers.Create("jrbTestServer",sourceImage,n.Flavors.List()[1],serverOptions)
    for s.Status != "ACTIVE" && s.Status != "ERROR" {
        s = n.Servers.Get(s.Id)
    }
    fmt.Println("Server has been created.  Attaching volume...")
    attachOptions := make(map[string]interface{})
    cdr.Volumes.Attach(v.Id,s.Id,"/dev/vdb",attachOptions)
    fmt.Println("DONE!")
}
