package volumes

import (
    "fmt"
    "os"
    "go_openstack_client/authhttp/authenticator"
    "go_openstack_client/authhttp/none"
    "go_openstack_client/apiconnection"
    "go_openstack_client/testserver"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type VolumesTestSuite struct{
    TestServer testserver.TestServer
}

var _ = gocheck.Suite(&VolumesTestSuite{})

func (t *VolumesTestSuite) SetUpSuite (c *gocheck.C) {
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


func (t *VolumesTestSuite) xTest_List (c *gocheck.C) {
    //apiConn := apiconnection.New("http://127.0.0.1:" + t.TestServer.Port,"bosh","bosh","bosh")
    apiConn := apiconnection.New("http://10.150.0.60:35357","volume","bosh","bosh","bosh")
    volumes := New(apiConn)
    volumeList := volumes.List()
    fmt.Println(volumeList)
}

func (t *VolumesTestSuite) xTest_Get (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","volume","bosh","bosh","bosh")
    volumes := New(apiConn)
    v := volumes.Get("56bd7e8b-eee3-4624-a31d-e0f22e3eabf0")
    fmt.Println(v)
}


func (t *VolumesTestSuite) xTest_Create (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","volume","bosh","bosh","bosh")
    volumes := New(apiConn)
    options := make(map[string]interface{})
    v := volumes.Create("jrbNewVolume",float64(20),options)
    fmt.Println(v)
    fmt.Println(v.Id)
}

func (t *VolumesTestSuite) xTest_Detach (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","volume","bosh","bosh","bosh")
    volumes := New(apiConn)
    volumes.Detach("bc342444-9f0f-495d-b711-bf96b55bf57d")
}

func (t *VolumesTestSuite) Test_Delete (c *gocheck.C) {
    apiConn := apiconnection.New("http://10.150.0.60:35357","volume","bosh","bosh","bosh")
    volumes := New(apiConn)
    volumes.Delete("48ffcb93-3b47-4b3e-817a-8ba80174b11c")
}
