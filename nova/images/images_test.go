package images

import (
    "fmt"
    "os"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/none"
    "go-openstack-client/apiconnection"
    "go-openstack-client/testserver"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ImagesTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&ImagesTestSuite{})

func (t *ImagesTestSuite) SetUpSuite (c *gocheck.C) {
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


func (t *ImagesTestSuite) Test_List (c *gocheck.C) {
    //apiConn := apiconnection.New("http://127.0.0.1:" + t.TestServer.Port,"bosh","bosh","bosh")
    apiConn := apiconnection.New("http://10.150.0.60:35357","nova","bosh","bosh","bosh")
    images := New(apiConn)
    fmt.Println(images.List())
}