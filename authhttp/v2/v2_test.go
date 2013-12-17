package v2

import (
    "fmt"
    "launchpad.net/gocheck"
    "testing"
    "os"
    "git.smf.sh/jrbudnack/go_openstack_client/authhttp/authenticator"
    "git.smf.sh/jrbudnack/go_openstack_client/authhttp/client"
    "git.smf.sh/jrbudnack/go_openstack_client/authhttp/none"
    "git.smf.sh/jrbudnack/go_openstack_client/testserver"
    "git.smf.sh/jrbudnack/go_openstack_client/servicecatalog"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type V2TestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&V2TestSuite{})

func (t *V2TestSuite) SetUpSuite (c *gocheck.C) {
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
    endpointQuery["servicename"] = "compute"
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
    testClient := client.New(testCreds, "http://127.0.0.1:" + t.TestServer.Port)
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
    testClient := client.New(testCreds, "http://127.0.0.1:" + t.TestServer.Port)
    retval, _ := testClient.PostAndParseBody("/","")

    return testCreds, string(retval)
}
