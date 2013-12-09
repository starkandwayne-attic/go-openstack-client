package none

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/client"
    "go-openstack-client/authhttp/mockserver"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type NoAuthTestSuite struct{}

var _ = gocheck.Suite(&NoAuthTestSuite{})

func (t *NoAuthTestSuite) SetUpSuite (c *gocheck.C) {
    authorizers := authenticator.Authenticators{}
    authorizers.Add(Authenticator{},true)

    mockServer := mockserver.Server{}
    go mockServer.Start(authorizers, "8082")
}

func (t *NoAuthTestSuite) Test_Authorization(c *gocheck.C) {
    testClient := client.New(Credentials{}, "http://localhost:8082")
    retval, _ := testClient.GetAndParseBody("/")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}
