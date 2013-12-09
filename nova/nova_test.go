package nova

import (
    _"fmt"
    _"time"
    "os"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/none"
    "go-openstack-client/testserver"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type NovaTestSuite struct{}

var _ = gocheck.Suite(&NovaTestSuite{})

func (t *NovaTestSuite) SetUpSuite (c *gocheck.C) {
    authenticators := authenticator.Authenticators{}
    // We use Authentication = none because we aren't writing the 
    // authentication server.  We are only writing the client.
    // Therefore, we only need the API server to simulate responses,
    // not do actual authentication.
    authenticators.Add(none.Authenticator{},true)

    testServer := testserver.TestServer{}
    workingDir, _ := os.Getwd()
    go testServer.Start(authenticators, "8083", workingDir + "/testfiles")
}


func (t *NovaTestSuite) Test_Authorization_NoCreds (c *gocheck.C) {
    //time.Sleep(30 * time.Second)
    n := New("http://127.0.0.1:8083","bosh","bosh","bosh")
    n.PrintServerList()
}
