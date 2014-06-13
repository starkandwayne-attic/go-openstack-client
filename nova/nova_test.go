package nova

import (
	"fmt"
	"github.com/starkandwayne/go-openstack-client/apitestharness"
	"github.com/starkandwayne/go-openstack-client/cinder"
	"github.com/starkandwayne/go-openstack-client/quantum"
	"launchpad.net/gocheck"
    "strconv"
	"testing"
    "sync"
    "time"
    "math/rand"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ServersTestSuite struct {
	NovaApiTestHarness    apitestharness.ApiTestHarness
	CinderApiTestHarness  apitestharness.ApiTestHarness
	QuantumApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&ServersTestSuite{})

func (t *ServersTestSuite) SetUpSuite(c *gocheck.C) {
	t.NovaApiTestHarness = apitestharness.New("compute", false)
	t.CinderApiTestHarness = apitestharness.New("volume", false)
	t.QuantumApiTestHarness = apitestharness.New("network", false)
}

func (t *ServersTestSuite) xTest_ListIPS(c *gocheck.C) {
	n := New(t.NovaApiTestHarness.Url,
		t.NovaApiTestHarness.Username,
		t.NovaApiTestHarness.Password,
		t.NovaApiTestHarness.Tenant)
    fips, _ := n.FloatingIps.List()

    for _, floatingIp := range fips {
        println(floatingIp.Id, " ", floatingIp.Ip)
    }
}

func (t *ServersTestSuite) xTest_ReserveFloatingIP(c *gocheck.C) {
	n := New(t.NovaApiTestHarness.Url,
		t.NovaApiTestHarness.Username,
		t.NovaApiTestHarness.Password,
		t.NovaApiTestHarness.Tenant)
    fip, err := n.FloatingIps.Create()
    if err != nil {
        println(err.Error())
    }
    println("FIP ID: ", fip.Id, " FIP: ", fip.Ip)
}

func (t *ServersTestSuite) xTest_AssignZeroFIP(c *gocheck.C) {
	n := New(t.NovaApiTestHarness.Url,
		t.NovaApiTestHarness.Username,
		t.NovaApiTestHarness.Password,
		t.NovaApiTestHarness.Tenant)
    err := n.FloatingIps.AttachToServer("4770e2ab-b96e-4b32-8c21-da070f5873a2", "0")
    if err != nil {
        println(err.Error())
    }
}

func (t *ServersTestSuite) Test_CreateServers(c *gocheck.C) {
	n := New(t.NovaApiTestHarness.Url,
		t.NovaApiTestHarness.Username,
		t.NovaApiTestHarness.Password,
		t.NovaApiTestHarness.Tenant)
	qtm := quantum.New(t.QuantumApiTestHarness.Url,
		t.QuantumApiTestHarness.Username,
		t.QuantumApiTestHarness.Password,
		t.QuantumApiTestHarness.Tenant)
	serverOptions := make(map[string]interface{})
	serverOptions["keyname"] = "jrb_snw"

	availableNetwork, _ := qtm.Networks.GetByName("int-net")

	privateNet := make(map[string]string)
	privateNet["uuid"] = availableNetwork.Id

	networksList := make([]map[string]string, 0)
	networksList = append(networksList, privateNet)

	serverOptions["networks"] = networksList

	sourceImage, _ := n.Images.GetByName("ubuntu12.04-no-metaservice-sshd")
	sourceFlavor, _ := n.Flavors.GetByName("m1.small")

    fips, _ := n.FloatingIps.List()

    i := 0
    var wg sync.WaitGroup
    for _, floatingIp := range fips {
        i++
        wg.Add(1)
        go func(procnum int, fip string, fipid string) {
            time.Sleep(time.Duration(rand.Intn(10)) * time.Nanosecond)
            println("Creating server for: ", fipid, " ", fip)
        	s, _ := n.Servers.Create("fip_test_api_" + strconv.Itoa(procnum) , sourceImage, sourceFlavor, serverOptions)
            println("Waiting for server...")
        	for s.Status != "ACTIVE" && s.Status != "ERROR" {
	        	s, _ = n.Servers.Get(s.Id)
    	    }
    	    fmt.Println("Attaching floating ip...")
        	n.FloatingIps.AttachToServer(s.Id, fipid)
            wg.Done()
        }(i,floatingIp.Ip, floatingIp.Id)
    }
    wg.Wait()
}

func (t *ServersTestSuite) xTest_CreateAServerEndToEnd(c *gocheck.C) {
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
	volumeOptions := make(map[string]interface{})
	v, _ := cdr.Volumes.Create("jrbNewVolume", float64(20), volumeOptions)
    fmt.Println("Waiting for volume " + v.DisplayName + " to become available.")
	for v.Status != "available" && v.Status != "error" {
		v, _ = cdr.Volumes.Get(v.Id)
	}
    fmt.Println("Volume has been created.  Creating server...")
	serverOptions := make(map[string]interface{})
	serverOptions["keyname"] = "jrb_snw"

	availableNetwork, _ := qtm.Networks.GetByName("int-net")

	privateNet := make(map[string]string)
	privateNet["uuid"] = availableNetwork.Id

	networksList := make([]map[string]string, 0)
	networksList = append(networksList, privateNet)

	serverOptions["networks"] = networksList

	sourceImage, _ := n.Images.GetByName("ubuntu12.04-no-metaservice-sshd")
	sourceFlavor, _ := n.Flavors.GetByName("m1.small")
	s, _ := n.Servers.Create("jrbTestServer3", sourceImage, sourceFlavor, serverOptions)
	for s.Status != "ACTIVE" && s.Status != "ERROR" {
		s, _ = n.Servers.Get(s.Id)
	}
	fmt.Println("Server has been created.  Attaching volume...")
	attachOptions := make(map[string]interface{})
	cdr.Volumes.Attach(v.Id, s.Id, "/dev/vdb", attachOptions)
	fmt.Println("Attaching floating ip...")
	n.FloatingIps.CreateAndAttachToServer(s.Id)
	fmt.Println("DONE!")
}
