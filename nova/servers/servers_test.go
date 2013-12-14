package servers

import (
    "fmt"
    "os"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/none"
    "go-openstack-client/apiconnection"
    "go-openstack-client/testserver"
    "go-openstack-client/nova/images"
    "go-openstack-client/nova/flavors"
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


func (t *ServersTestSuite) Test_List (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","nova","bosh","bosh","bosh")
    servers := New(apiConn)
    fmt.Println(servers.List())
}

func (t *ServersTestSuite) Test_Get (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","nova","bosh","bosh","bosh")
    servers := New(apiConn)
    s := servers.Get("c4c1630b-de71-44b2-aea4-e52067e149fb")
    fmt.Println(s)
}

func (t *ServersTestSuite) xTest_Create(c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","nova","bosh","bosh","bosh")
    servers := New(apiConn)
    images := images.New(apiConn)
    flavors := flavors.New(apiConn)
    options := make(map[string]interface{})

    servers.Create("jrbTestServer",images.List()[0],flavors.List()[1],options)
}
