package v2

import (
    "fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/client"
    "go-openstack-client/authhttp/none"
    "go-openstack-client/testserver"
    "go-openstack-client/servicecatalog"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type V2TestSuite struct{}

var _ = gocheck.Suite(&V2TestSuite{})

func (t *V2TestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    authenticators.Add(none.Authenticator{},true)

    testServer := testserver.TestServer{}
    go testServer.Start(authenticators, "8082", "/data/Projects/go/src/go-openstack-client/authhttp/v2/testfiles")
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
    creds, retval := t.testUserCreds("glance","servicepass")
    endpointQuery := make(map[string]string)
    endpointQuery["urltype"] = "public"
    endpointQuery["servicename"] = "nova"
    sc := creds["serviceCatalog"].(servicecatalog.ServiceCatalog)
    fmt.Println(string(retval))
    //NOTE:  Feed this token from the results in the first pass
    creds, retval = t.testTokenCreds(creds["token"].(string))
    fmt.Println(string(retval))
    fmt.Println(sc.GetEndpoint(endpointQuery))
    //c.Assert(string(retval), gocheck.Equals, "<html><body><h1>GET Successful!</h1></body></html>")
}

func (t *V2TestSuite) testTokenCreds (token string) (Credentials, string) {
    testCreds := Credentials{}
    if token != "" {
        testCreds["token"] = token
    }
    testClient := client.New(testCreds, "http://127.0.0.1:8082")
    retval, _ := testClient.PostAndParseBody("/","")

    return testCreds, string(retval)
}

func (t *V2TestSuite) testUserCreds (user string, password string) (Credentials, string) {
    testCreds := Credentials{}
    if user != "" {
        testCreds["username"] = user
        testCreds["password"] = password
        testCreds["tenantName"] = "service"
    }
    testClient := client.New(testCreds, "http://127.0.0.1:8082")
    retval, _ := testClient.PostAndParseBody("/","")

    return testCreds, string(retval)
}
