package basic

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

type BasicAuthTestSuite struct{}

func (t *BasicAuthTestSuite) users() map[string]string {
  users := make(map[string]string)
  users["testuser"] = "dGVzdHVzZXI6cGlja2xlcw==" //password is "pickles"
  return users
}

var _ = gocheck.Suite(&BasicAuthTestSuite{})

func (t *BasicAuthTestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    authenticators.Add(Authenticator{users: t.users},true)

    mockServer := mockserver.Server{}
    go mockServer.Start(authenticators, "8081")
}

func (t *BasicAuthTestSuite) Test_Authorization_NoCreds (c *gocheck.C) {
    retval := t.testCreds("","")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Valid Credentials Not Supplied</h1></body></html>")
}

func (t *BasicAuthTestSuite) Test_Authorization_InvalidCreds (c *gocheck.C) {
    retval := t.testCreds("invaliduser","fffffffffffffffffffffffffffffff")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Could Not Find User</h1></body></html>")
}

func (t *BasicAuthTestSuite) Test_Authorization_WrongCreds (c *gocheck.C) {
    retval := t.testCreds("testuser","fffffffffffffffffffffffffffffff")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Password Mismatch</h1></body></html>")
}

func (t *BasicAuthTestSuite) Test_Authorization_Successful (c *gocheck.C) {
    retval := t.testCreds("testuser","pickles")
    c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}

func (t *BasicAuthTestSuite) testCreds (user string, password string) string {
    testCreds := Credentials{}
    if user != "" {
        testCreds["username"] = user
        testCreds["password"] = password
    }
    testClient := client.New(testCreds, "http://localhost:8081")
    retval, _ := testClient.GetAndParseBody("/")

    return string(retval)
}
