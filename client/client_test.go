package client

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/mockauthentication"
    "go-openstack-client/testserver"
)

// HoRk up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ClientTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&ClientTestSuite{})

func (t *ClientTestSuite) SetUpSuite (c *gocheck.C) {
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

