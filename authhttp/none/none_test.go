package none

import (
    _"fmt"
    "launchpad.net/gocheck"
    "os"
    "testing"
    "git.smf.sh/jrbudnack/go_openstack_client/testserver"
    "git.smf.sh/jrbudnack/go_openstack_client/authhttp/authenticator"
    "git.smf.sh/jrbudnack/go_openstack_client/authhttp/client"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type NoAuthTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&NoAuthTestSuite{})

func (t *NoAuthTestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    authenticators.Add(Authenticator{},true)

    workingDir, _ := os.Getwd()
    t.TestServer = testserver.New(authenticators, "", workingDir + "/testfiles")
    go t.TestServer.Start()
}

func (t *NoAuthTestSuite) Test_Authorization(c *gocheck.C) {
    testClient := client.New(Credentials{}, "http://localhost:" + t.TestServer.Port)
    retval, _ := testClient.GetAndParseBody("/")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}
