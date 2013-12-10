package apiconnection

import (
    _"fmt"
    "os"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/none"
    "go-openstack-client/testserver"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ApiConnectionTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&ApiConnectionTestSuite{})

func (t *ApiConnectionTestSuite) SetUpSuite (c *gocheck.C) {
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


func (t *ApiConnectionTestSuite) Test_GetServerList (c *gocheck.C) {
    n := New("http://127.0.0.1:" + t.TestServer.Port,"bosh","bosh","bosh")
    n.PrintServerList()
}
