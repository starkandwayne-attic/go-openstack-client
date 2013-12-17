package apiconnection

import (
    _"fmt"
    "os"
    "go_openstack_client/authhttp/authenticator"
    "go_openstack_client/authhttp/none"
    "go_openstack_client/testserver"
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
    t.RunTestApiServer()
}

func (t *ApiConnectionTestSuite) RunTestApiServer() {
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
