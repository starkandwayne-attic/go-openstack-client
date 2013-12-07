package client

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/mockauthentication"
    "go-openstack-client/authhttp/mockserver"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ClientTestSuite struct{}

var _ = gocheck.Suite(&ClientTestSuite{})

func (t *ClientTestSuite) SetUpSuite (c *gocheck.C) {
    authorizers := authenticator.Authenticators{}
    authorizers.Add(mockauthentication.Authenticator{},true)

    mockServer := mockserver.Server{}
    go mockServer.Start(authorizers, "8083")
}

func (t *ClientTestSuite) Test_GetAndParseBody_NoCreds (c *gocheck.C) {
    client := t.prepareClientNoCreds()
    retval, _ := client.GetAndParseBody("/")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Unauthorized</h1></body></html>")
}

func (t *ClientTestSuite) Test_GetAndParseBody_WithCreds (c *gocheck.C) {
    client := t.prepareClientWithCreds()
    retval, _ := client.GetAndParseBody("/")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}

func (t *ClientTestSuite) Test_PostAndParseBody_NoCreds (c *gocheck.C) {
    client := t.prepareClientNoCreds()
    retval, _ := client.PostAndParseBody("/","Test")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Unauthorized</h1></body></html>")
}

func (t *ClientTestSuite) Test_PostAndParseBody_WithCreds (c *gocheck.C) {
    client := t.prepareClientWithCreds()
    retval, _ := client.PostAndParseBody("/","Test")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>POST Successful!</h1><h2>Test</h2></body></html>")
}

func (t *ClientTestSuite) prepareClientWithCreds() Client {
    creds := mockauthentication.Credentials{"user": "Jeremy", "password": "PicklesRGR8"}
    return New(creds, "http://localhost:8083")
}

func (t *ClientTestSuite) prepareClientNoCreds() Client {
    return New(mockauthentication.Credentials{}, "http://localhost:8083")
}
