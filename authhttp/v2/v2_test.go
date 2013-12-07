package v2

import (
    "fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/client"
    "go-openstack-client/authhttp/mockserver"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type V2TestSuite struct{}

func (t *V2TestSuite) users() map[string]string {
  users := make(map[string]string)
  users["testuser"] = "dGVzdHVzZXI6cGlja2xlcw==" //password is "pickles"
  return users
}

var _ = gocheck.Suite(&V2TestSuite{})

func (t *V2TestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    authenticators.Add(Authenticator{users: t.users},true)

    mockServer := mockserver.Server{}
    go mockServer.Start(authenticators, "8082")
}

func (t *V2TestSuite) Test_Authorization_NoCreds (c *gocheck.C) {
    //retval := t.testCreds("","")
    //c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Valid Credentials Not Supplied</h1></body></html>")
}

func (t *V2TestSuite) Test_Authorization_InvalidCreds (c *gocheck.C) {
    //retval := t.testCreds("invaliduser","fffffffffffffffffffffffffffffff")
    //c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Could Not Find User</h1></body></html>")
}

func (t *V2TestSuite) Test_Authorization_WrongCreds (c *gocheck.C) {
    //retval := t.testCreds("testuser","fffffffffffffffffffffffffffffff")
    //c.Assert(string(retval), gocheck.Equals, "<html><body><h1>Password Mismatch</h1></body></html>")
}

func (t *V2TestSuite) Test_Authorization_Successful (c *gocheck.C) {
    retval := t.testUserCreds("glance","servicepass")
    fmt.Println(string(retval))
    //NOTE:  Feed this token from the results in the first pass
    retval = t.testTokenCreds("d819998ca5974007b71aad6fe187b7a0")
    fmt.Println(string(retval))

    //c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}

func (t *V2TestSuite) testTokenCreds (token string) string {
    testCreds := Credentials{}
    if token != "" {
        testCreds["token"] = token
    }
    testClient := client.New(testCreds, "http://10.150.0.60:35357")
    retval, _ := testClient.PostAndParseBody("/v2.0/tokens","")

    return string(retval)
}

func (t *V2TestSuite) testUserCreds (user string, password string) string {
    testCreds := Credentials{}
    if user != "" {
        testCreds["username"] = user
        testCreds["password"] = password
        testCreds["tenantName"] = "service"
    }
    testClient := client.New(testCreds, "http://10.150.0.60:35357")
    retval, _ := testClient.PostAndParseBody("/v2.0/tokens","")

    return string(retval)
}
