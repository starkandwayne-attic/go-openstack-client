package nova

import (
    "fmt"
    "git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
    "git.smf.sh/jrbudnack/go_openstack_client/cinder"
    "git.smf.sh/jrbudnack/go_openstack_client/quantum"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ServersTestSuite struct{
    NovaApiTestHarness apitestharness.ApiTestHarness
    CinderApiTestHarness apitestharness.ApiTestHarness
    QuantumApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&ServersTestSuite{})

func (t *ServersTestSuite) SetUpSuite (c *gocheck.C) {
    t.NovaApiTestHarness = apitestharness.New("compute", false)
    t.CinderApiTestHarness = apitestharness.New("volume", false)
    t.QuantumApiTestHarness = apitestharness.New("network", false)
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
    qtm := quantum.New(t.QuantumApiTestHarness.Url,
                       t.QuantumApiTestHarness.Username,
                       t.QuantumApiTestHarness.Password,
                       t.QuantumApiTestHarness.Tenant)
    fmt.Println("Volume has been created.  Creating server...")
    volumeOptions := make(map[string]interface{})
    v, _ := cdr.Volumes.Create("jrbNewVolume",float64(20),volumeOptions)
    for v.Status != "available" && v.Status != "error" {
        v, _ = cdr.Volumes.Get(v.Id)
    }
    serverOptions := make(map[string]interface{})
    serverOptions["keyname"] = "bosh"

    availableNetwork, _ := qtm.Networks.GetByName("cloud")

    privateNet := make(map[string]string)
    privateNet["uuid"] = availableNetwork.Id

    networksList := make([]map[string]string,0)
    networksList = append(networksList, privateNet)

    serverOptions["networks"] = networksList

    sourceImage, _ := n.Images.GetByName("centos")
    sourceFlavor, _ := n.Flavors.GetByName("m1.small")
    s, _ := n.Servers.Create("jrbTestServer",sourceImage,sourceFlavor,serverOptions)
    for s.Status != "ACTIVE" && s.Status != "ERROR" {
        s, _ = n.Servers.Get(s.Id)
    }
    fmt.Println("Server has been created.  Attaching volume...")
    attachOptions := make(map[string]interface{})
    cdr.Volumes.Attach(v.Id,s.Id,"/dev/vdb",attachOptions)
    fmt.Println("DONE!")
}
