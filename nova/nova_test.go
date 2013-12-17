package nova

import (
    "fmt"
    "os"
    "go_openstack_client/authhttp/authenticator"
    "go_openstack_client/authhttp/none"
    "go_openstack_client/testserver"
    "go_openstack_client/cinder"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ServersTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&ServersTestSuite{})

func (t *ServersTestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    // We use Authentication = none because we aren't writing the 
    // authentication server.  We are only writing the client.
    // Therefore, we only need the API server to simulate responses,
    // not do actual authentication.
    authenticators.Add(none.Authenticator{},true)

    workingDir, _ := os.Getwd()
    t.TestServer = testserver.New(authenticators, "", workingDir + "/testfiles")
    go t.TestServer.Start()
}

func (t *ServersTestSuite) Test_CreateAFreakingServer (c *gocheck.C) {
    fmt.Println("Creating volume...")
    n := New("http://172.17.74.3:5000","cf_test3","cf_test3","cf_test3")
    cdr := cinder.New("http://172.17.74.5000","cf_test3","cf_test3","cf_test3")

    fmt.Println("Volume has been created.  Creating server...")
    volumeOptions := make(map[string]interface{})
    v := cdr.Volumes.Create("jrbNewVolume",float64(20),volumeOptions)
    for v.Status != "available" && v.Status != "error" {
        v = cdr.Volumes.Get(v.Id)
    }
    serverOptions := make(map[string]interface{})
    s := n.Servers.Create("jrbTestServer",n.Images.List()[0],n.Flavors.List()[1],serverOptions)
    for s.Status != "ACTIVE" && s.Status != "ERROR" {
        s = n.Servers.Get(s.Id)
    }
    fmt.Println("Server has been created.  Attaching volume...")
    attachOptions := make(map[string]interface{})
    cdr.Volumes.Attach(v.Id,s.Id,"/dev/vdb",attachOptions)
    fmt.Println("DONE!")
}
